<template>
  <div class="message" v-if="msg && chats">
    <div v-if="msg.forwarded" class="forwarded-label">>> Forwarded</div>
    <div v-if="!msg.forwarded && msg.reply !== 0" class="reply-preview" style="background-color:blanchedalmond">
      <div class="row">
        <p class="text-left ms-2 mt-2">Replied to</p>
      </div>
      <div class="row">
        <p class="text-left ms-2 mt-2">{{ messages.find(m => m.message_id == msg.reply)?.content ?? "Deleted message"
          }}
        </p>
      </div>
    </div>

    <div class="message-content">
      <img id="photo" :src="source + '/chats/' + msg.chat_id + '/media/' + msg.photo_id">
    </div>

    <div class="reactions mt-4">
      <button v-for="(reaction, index) in reactions" :key="reaction" @click="addReaction(reaction)"
        :class="{ 'selected-reaction': selectedReaction && reaction == selectedReaction.reaction && msg.message_id == selectedReaction.message_id, 'col-3 btn btn-light btn-sm': true }">
        {{ emoji[index] }}
      </button>
    </div>
    <div class="mt-1 hover-box" v-if="reactionsUsers" @click="showReactionsPopup = !showReactionsPopup">
      <p style="font-style: italic;">Total reactions: {{ reactionsUsers.length }}</p>
    </div>

    <!-- Popup for listing reactions -->
    <div v-if="showReactionsPopup" class="popup-overlay">
      <div class="popup-content">
        <h6>Reactions</h6>

        <div class="row" v-for="reaction in reactionsUsers" :key="reaction.user">
          <div class="col-12"> <strong>{{ reaction.owner }}:</strong> {{ emoji[reaction.reaction] }} </div>
        </div>

        <button class="btn btn-danger" @click="showReactionsPopup = false">Close</button>
      </div>
    </div>
    <div class="row mt-2 mb-2"> <!--v-if="!msg.forwarded && msg.reply === 0"-->
      <div class="col-4">
        <button class="btn btn-info" style="color: white;" @click="showReplyPopup">Reply</button>
      </div>
      <div class="col-4">
        <button class="btn btn-success" @click="showForwardPopup">Forward</button>
      </div>
      <div class="col-4" v-if="this.identifier == this.msg.user_id">
        <button class="btn btn-danger" @click="deleteMessage">Delete</button>
      </div>
    </div>

    <div class="message-info text-center">
      <span class="user-id me-2" v-if="!msg.forwarded">{{ msg.user_id }}</span>
      <span class="message-date">{{ formatDate(msg.date) }}</span>
    </div>
    <div class="row mt-2" v-if=" msg.user_id == identifier">
      <p class="text-end" style="font-style: italic;">{{ !msg.status ? 'Sent v' : 'Read w' }} </p>
    </div>

    <div v-if="showPopup2" class="popup">
      <div class="popup-content">
        <h3>Reply to this message</h3>
        <div class="row">
          <div class="col-12">
            <input type="text" v-model="text" class="form-control" placeholder="Type a message..." />
          </div>
          <div class="col-12">
            <button class="btn btn-info mt-2" :disabled="this.text.length == 0"
              @click="replyToMessage(msg.message_id)">Send reply</button>
          </div>
        </div>
        <button class="btn btn-danger mt-2" @click="closePopup2">Close</button>
      </div>
    </div>

    <div v-if="showPopup" class="popup">
      <div class="popup-content">
        <h6 class="font-weight-bold">Select a user to forward</h6>
          <form class="d-flex mb-3" role="search">
            <input class="form-control me-2" type="search" placeholder="@username" v-model="userToForward">
                <button class="btn btn-outline-success" type="submit"
                  :disabled="!this.userToForward || this.userToForward.trim().length == 0"
                  @click="getUsers">Search</button>
          </form>
          <div class="row" v-for="(user, index) in foundUsers" :key="index">
              <div class="col-6">
                {{ user.user_id}}
              </div>
              <div class="col-6">
                <button class="btn btn-info" @click="forwardMessageToUser(user.user_id)">Forward</button>
              </div>
          </div>
        <h6 class="font-weight-bold">Select a chat to forward</h6>
        <div class="row" v-for="chat in chats" :key="chat.chat_id">
          <div class="col-7">
            {{ chat.Chat_name }}
          </div>
          <div class="col-5">
            <button class="btn btn-info" @click="forwardMessage(chat.first_chat_id)">Forward</button>
          </div>
        </div>
        <button class="btn btn-danger" @click="closePopup">Close</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ["msg", "messages", "identifier", "chats"],
  data() {
    return {
      emoji: ["<3", ":)", ">:|", ":("],
      reactions: [0, 1, 2, 3],
      selectedReaction: null,
      forwadedToChat: null,
      text: "",
      showPopup: false,
      showPopup2: false,
      reactionsUsers: [],
      showReactionsPopup: false,
      userToForward: "",
      foundUsers : null,
      source: __API_URL__
    };
  },

  async mounted() {
    await this.getReactions();

  },

  methods: {
    formatDate(dateString) {
      const options = { year: 'numeric', month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' };
      const date = new Date(dateString);
      return date.toLocaleDateString('en-US', options);
    },
    showForwardPopup() {
      this.showPopup = true;
    },
    closePopup() {
      this.showPopup = false;
    },
    showReplyPopup() {
      this.showPopup2 = true;
    },
    closePopup2() {
      this.showPopup2 = false;
    },

    async getReactions() {
      try {
        const response = await this.$axios.get(`chats/${this.msg.chat_id}/messages/${this.msg.message_id}/reaction`);
        this.reactionsUsers = response.data;

        let r = null
        if (this.reactionsUsers) { r = this.reactionsUsers.filter(u => u.owner == this.identifier) }
        if (r) { this.selectedReaction = r[0] }
      } catch (error) {
        console.log(error);
      }
    },

    async getUsers() {
      try {
            let response = await this.$axios.get("/users", {
              params: {
                 user_id: this.userToForward.trim()
              }
                });
                this.foundUsers = response.data
                this.userToForward = ""
            }
            catch (error) {
                console.log(error)
            }
    },

    async addReaction(reaction) {
      try {

        if (this.selectedReaction && reaction == this.selectedReaction.reaction) {
          await this.$axios.delete(`chats/${this.msg.chat_id}/messages/${this.msg.message_id}/reaction/${reaction}`);
          this.selectedReaction = null;
        } else if (this.selectedReaction && reaction != this.selectedReaction.reaction) {
          await this.$axios.delete(`chats/${this.msg.chat_id}/messages/${this.msg.message_id}/reaction/${this.selectedReaction.reaction}`);
          await this.$axios.put(`chats/${this.msg.chat_id}/messages/${this.msg.message_id}/reaction/${reaction}`);
        } else {
          await this.$axios.put(`chats/${this.msg.chat_id}/messages/${this.msg.message_id}/reaction/${reaction}`);
        }

        await this.getReactions()
        this.$emit("newMessage")
      } catch (error) {
        console.log(error);
      }
    },

    async forwardMessage(chatId) {
      try {
        await this.$axios.post("/chats/" + chatId + "/forwarded/" + this.msg.message_id, {
          chat_id: this.msg.chat_id.toString()
        });
        this.showPopup = false;
        this.$emit("newMessage")
        // window.location.reload(); 
      } catch (error) {
        console.log(error);
      }
    },

    async forwardMessageToUser(user_id) {
      try {
        var trovato = false
        for(const chat of this.chats) {
          let cn = chat.Chat_name.split("-")
          if(cn.length>1 && (cn[0] == user_id || cn[1] == user_id)) {
            console.log("qui--->",chat)
            this.forwardMessage(chat.first_chat_id)
            trovato = true
            break
          }
        }
        if(!trovato) {
          let response = await this.$axios.put("/chats", {
                    user_ids: [this.identifier.toString(), user_id.toString()],
                    chat_group: false
                }
                );
            this.forwardMessage(response.data.chat_id)
        }
      }
          

      catch(error) {
        console.log(error)
      }
    },

    async deleteMessage() {
      try {
        await this.$axios.delete("/chats/" + this.msg.chat_id + "/messages/" + this.msg.message_id);
        window.location.reload(); 

      }
      catch (error) {
        console.log(error)
      }
    },

    async replyToMessage(msg_id) {

      try {
        await this.$axios.post("/chats/" + this.msg.chat_id + "/messages", {
          chat_id: this.msg.chat_id,
          content: this.text,
          forwarded: 0,
          reply: parseInt(msg_id),
        });

        this.text = "";
        this.showPopup2 = false;
        this.$emit("newMessage")
        // window.location.reload(); 

      }
      catch (error) {
        console.log(error)
      }
    }
  }
}
</script>

<style scoped>
.message {
  padding: 10px;
  border-radius: 10px;
  margin: 10px 0;
}

.message.regular {
  background-color: #e5e5e5;
}

.forwarded-label {
  font-weight: bold;
  color: #555;
}

.message.forwarded {
  background-color: #dcf8c6;
}

.message.reply {
  background-color: #f0f0f0;
  margin-left: 20px;
}

.reply-preview {
  font-style: italic;
  color: #777;
  margin-bottom: 5px;
}

.message-info {
  font-size: 12px;
  color: #777;
}

.message-content {
  margin-top: 5px;
  font-size: 14px;
  color: #000;
}


.reactions {
  margin-top: 5px;
}

button {
  border: none;
  cursor: pointer;
  font-size: 14px;
  margin-right: 5px;
}

.hovered:hover {
  background-color: whitesmoke;
}

button.selected-reaction {
  background-color: yellow;
}

.popup {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: rgb(187, 187, 187);
  padding: 20px;
  box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
  z-index: 1000;
}

.popup-content {
  text-align: center;
}
</style>
