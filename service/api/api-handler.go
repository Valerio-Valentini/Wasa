package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// USER
	// rt.router.PUT("/users/:user_id", rt.putNewUserNickname)    // ok +
	rt.router.PUT("/users/:user_id/photo", rt.putNewUserPhoto) // notok +
	// SEARCH
	rt.router.GET("/users", rt.searchUser) // ok +
	// CHAT
	rt.router.GET("/users/:user_id/chats", rt.getAllChats)                  // ok +
	rt.router.GET("/users/:user_id/chats/:chat_id", rt.getMessagesFromChat) // ok +
	// GROUP
	rt.router.PUT("/chats", rt.createNewChat)                            // ok +
	rt.router.PUT("/chats/:chat_id/photo", rt.putNewGroupPhoto)          // notok +
	rt.router.PUT("/chats/:chat_id/members/:member_id", rt.addAMember)   // ok +
	rt.router.PUT("/chats/:chat_id", rt.putNewGroupName)                 // ok +
	rt.router.DELETE("/chats/:chat_id/members/:member_id", rt.leaveChat) // ok +
	// MESSAGE
	rt.router.POST("/chats/:chat_id/messages", rt.sendmessage)                                       // ok +
	rt.router.DELETE("/chats/:chat_id/messages/:message_id", rt.deleteMessage)                       // ok +
	rt.router.GET("/chats/:chat_id/messages/:message_id/reaction", rt.getMessageReactions)           // ok +
	rt.router.PUT("/chats/:chat_id/messages/:message_id/reaction/:reaction_id", rt.commentMessage)   // ok +
	rt.router.DELETE("/chats/:chat_id/messages/:message_id/reaction/:reaction_id", rt.deleteComment) // ok +
	rt.router.POST("/chats/:chat_id/forwarded/:message_id", rt.forwardMessage)                       // ok +
	// MEDIA
	rt.router.POST("/chats/:chat_id/media/:media_id", rt.sendMedia)     // notok -
	rt.router.DELETE("/chats/:chat_id/media/:media_id", rt.deleteMedia) // notok -
	// LOGIN
	rt.router.POST("/session", rt.session) // ok +

	return rt.router
}
