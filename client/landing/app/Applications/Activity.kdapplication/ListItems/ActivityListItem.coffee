class ActivityListItemView extends KDListItemView

  getActivityChildConstructors = ->
    # CStatusActivity     : StatusActivityItemView
    JStatusUpdate       : StatusActivityItemView
    # CCodeSnipActivity   : CodesnipActivityItemView
    JCodeSnip           : CodesnipActivityItemView
    JQuestionActivity   : QuestionActivityItemView
    JDiscussionActivity : DiscussionActivityItemView
    JLinkActivity       : LinkActivityItemView

  getActivityChildCssClass = ->

    CFollowerBucket           : "system-message"
    CFolloweeBucket           : "system-message"
    CNewMemberBucket          : "system-message"
    CFollowerBucketActivity   : "system-message"
    CFolloweeBucketActivity   : "system-message"
    CNewMemberBucketActivity  : "system-message"

  getBucketMap =->
    JAccount  : AccountFollowBucketItemView
    JTag      : TagFollowBucketItemView

  constructor:(options = {},data)->
    
    options.type = "activity"
    
    super options, data
    
    {constructorName} = data.bongo_
    @setClass getActivityChildCssClass()[constructorName]

    unless options.isHidden
      if 'function' is typeof data.fetchTeaser
        data.fetchTeaser? (err, teaser)=> 
          @addChildView teaser
      else
        @addChildView data

  addChildView:(data, callback)->
    
    {constructorName} = data.bongo_

    childConstructor =
      if /CNewMemberBucket$/.test constructorName
        NewMemberBucketItemView
      else if /Bucket$/.test constructorName
        getBucketMap()[data.sourceName]
      else
        getActivityChildConstructors()[constructorName]

    if childConstructor
      childView = new childConstructor({}, data)
      @addSubView childView
      callback?()

  partial:-> ''

  show:->

    @getData().fetchTeaser? (err, teaser)=>
      @addChildView teaser, => @slideIn()

  slideIn:()-> @$().removeClass 'hidden-item'

    


class ActivityItemChild extends KDView

  constructor:(options, data)->

    origin =
      constructorName  : data.originType
      id               : data.originId

    @avatar = new AvatarView {
      size    : {width: 40, height: 40}
      origin
    }
    
    @author = new ProfileLinkView { origin }

    @tags = new ActivityChildViewTagGroup
      itemsToShow   : 3
      subItemClass  : TagLinkView
    , data.tags

    @commentBox = new CommentView null, data
    @actionLinks = new ActivityActionsView delegate : @commentBox.commentList, cssClass : "comment-header", data

    account = KD.whoami()
    if (data.originId is KD.whoami().getId()) or KD.checkFlag 'super-admin'
      @settingsButton = new KDButtonViewWithMenu
        cssClass    : 'transparent activity-settings-context activity-settings-menu'
        title       : ''
        icon        : yes
        delegate    : @
        iconClass   : "arrow"
        menu        : @settingsMenu data
        callback    : (event)=> @settingsButton.contextMenu event
    else
      @settingsButton = new KDCustomHTMLView tagName : 'span', cssClass : 'hidden'

    super
    
    data = @getData()
    data.on 'TagsChanged', (tagRefs)=>
      bongo.cacheable tagRefs, (err, tags)=>
        @getData().setAt 'tags', tags
        @tags.setData tags
        # debugger
        @tags.render()

    data.on 'PostIsDeleted', =>
      if KD.whoami().getId() is data.getAt('originId')
        @parent.destroy()
      else
        @parent.putOverlay
          isRemovable : no
          parent      : @parent
          cssClass    : 'half-white'

    data.watch 'repliesCount', (count)=>
      @commentBox.decorateCommentedState() if count >= 0

    @contentDisplayController = @getSingleton "contentDisplayController"

    bongo.cacheable origin.id, origin.constructorName, (err, account)->
      if account?.globalFlags and 'exempt' in account.globalFlags
        @setClass "exempt"

  settingsMenu:(data)->
    
    account = KD.whoami()
    
    menu = [
      type      : "contextmenu"
      items     : []
    ]

    if data.originId is KD.whoami().getId()
      menu[0].items = [
        { title : 'Edit',   id : 1,  parentId : null, callback : => @getSingleton('mainController').emit 'ActivityItemEditLinkClicked', data }
        { title : 'Delete', id : 2,  parentId : null, callback : => @confirmDeletePost data  }
      ]

      return menu
    
    if KD.checkFlag 'super-admin'
      menu[0].items = [
        { title : 'MARK USER AS TROLL', id : 1,  parentId : null, callback : => @markUserAsTroll data  }
        { title : 'UNMARK USER AS TROLL', id : 1,  parentId : null, callback : => @unmarkUserAsTroll data  }
        { title : 'Delete Post', id : 3,  parentId : null, callback : => @confirmDeletePost data  }
      ]

      return menu

  unmarkUserAsTroll:(data)->

    if data.originId
      bongo.cacheable "JAccount", data.originId, (err, account)->
        account.unflagAccount "exempt", (err, res)->
          if err then warn err
          else
            new KDNotificationView
              title : "@#{account.profile.nickname} won't be treated as a troll anymore!"

  markUserAsTroll:(data)->

    modal = new KDModalView
      title          : "MARK USER AS TROLL"
      content        : """
                        <div class='modalformline'>
                          This is what we call "Trolling the troll" mode.<br><br>
                          All of the troll's activity will disappear from the feeds, but the troll 
                          himself will think that people still gets his posts/comments.<br><br>
                          Are you sure you want to mark him as a troll? 
                        </div>
                       """
      height         : "auto"
      overlay        : yes
      buttons        :
        "YES, THIS USER IS DEFINITELY A TROLL" :
          style      : "modal-clean-red"
          loader     :
            color    : "#ffffff"
            diameter : 16
          callback   : =>
            # debugger
            if data.originId
              bongo.cacheable "JAccount", data.originId, (err, account)->
                account.flagAccount "exempt", (err, res)->
                  if err then warn err
                  else
                    modal.destroy()
                    new KDNotificationView
                      title : "@#{account.profile.nickname} marked as a troll!"


  confirmDeletePost:(data)->

    modal = new KDModalView
      title          : "Delete post"
      content        : "<div class='modalformline'>Are you sure you want to delete this post?</div>"
      height         : "auto"
      overlay        : yes
      buttons        :
        Delete       :
          style      : "modal-clean-red"
          loader     :
            color    : "#ffffff"
            diameter : 16
          callback   : =>
            data.delete (err)=>
              modal.buttons.Delete.hideLoader()
              modal.destroy()
              unless err then @emit 'ActivityIsDeleted'
              else new KDNotificationView
                type     : "mini"
                cssClass : "error editor"
                title     : "Error, please try again later!"
