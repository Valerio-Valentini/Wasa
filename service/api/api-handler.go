package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// USER
	rt.router.PUT("/users/:user_id", rt.putNewUserNickname) //ok
	rt.router.PUT("/users/:user_id/photo", rt.wrap(rt.putNewUserPhoto)) //ok
	//SEARCH
	rt.router.GET("/users", rt.wrap(rt.searchUser)) //ok
	//CHAT
	rt.router.GET("/users/:user_id/chats", rt.wrap(rt.getAllChats)) //ok
	rt.router.GET("/users/:user_id/chats/:chat_id", rt.wrap(rt.getMessagesFromChat)) //ok
	//GROUP
	rt.router.PUT("/users/:user_id/chats/:chat_id/photo", rt.wrap(rt.liveness)) //ok
	rt.router.PUT("/users/:user_id/chats/:chat_id/members/:member_id", rt.wrap(rt.addMember)) //ok
	rt.router.PUT("/users/:user_id/chats/:chat_id", rt.wrap(rt.liveness)) 
	rt.router.DELETE("/users/:user_id/chats/:chat_id", rt.wrap(rt.leaveChat)) //ok
	//MESSAGE
	rt.router.POST("/users/:user_id/chats/:chat_id/message/:message_id", rt.wrap(rt.sendmessage)) //ok
	rt.router.DELETE("/users/:user_id/chats/:chat_id/message/:message_id", rt.wrap(rt.deleteMessage)) //ok
	rt.router.PUT("/users/:user_id/chats/:chat_id/message/:message_id/reaction/:reaction_id", rt.wrap(rt.commentMessage)) //ok
	rt.router.DELETE("/users/:user_id/chats/:chat_id/message/:message_id/reaction/:reaction_id", rt.wrap(rt.deleteComment)) //ok
	rt.router.POST("/users/:user_id/chats/:chat_id/message/:message_id/forwarded/:message_id", rt.wrap(rt.ForwardMessage)) //ok
	//MEDIA
	rt.router.POST("/users/:user_id/chats/:chat_id/media/:media_id", rt.wrap(rt.sendMedia)) //ok
	rt.router.DELETE("/users/:user_id/chats/:chat_id/media/:media_id", rt.wrap(rt.deleteMedia)) //ok
	//LOGIN
	rt.router.POST("/session", rt.wrap(rt.session)) //ok
	// Special routes
	rt.router.GET("/liveness", rt.liveness) //ok

	return rt.router
}
