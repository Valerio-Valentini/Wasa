<template>
  <div class="message" v-if="msg">
    <div v-if="msg.forwarded" class="forwarded-label">>> Forwarded</div>
    <div v-if="msg.reply !== 0" class="reply-preview">
      {{ messages.find(m => m.message_id === msg.reply) || "Replied message not found" }}
    </div>
    <div class="message-info" v-if="!msg.forwarded">
      <span class="user-id me-2">{{ msg.user_id }}</span>
      <span class="message-date">{{ formatDate(msg.date) }}</span>
    </div>
    <div class="message-content">
      {{ msg.content }}
    </div>

    <div class="reactions">
      <button v-for="(reaction, index) in reactions" :key="reaction" @click="addReaction(reaction)"
        :class="{ 'selected-reaction': reaction === selectedReaction, 'hovered': true }">
        {{ emoji[index] }}
      </button>
    </div>
    <div class="message-actions" v-if="!msg.forwarded && msg.reply === 0">
      <button class="btn btn-info" style="color: white;" @click="replyToMessage">Reply</button>
      <button class="btn btn-success" @click="forwardMessage">Forward</button>
    </div>
    <div class="div">
      <button class="btn btn-danger" @click="deleteMessage">Delete message</button>
    </div>
  </div>
</template>

<script>
export default {
  props: ["msg", "messages", "identifier"],
  data() {
    return {
      emoji: ["<3", ":)", ">:(", ":("],
      reactions: [0, 1, 2, 3],
      selectedReaction: -1,
      forwadedToChat: null,
      text: ""
    };
  },

  mounted() {
    //load reaction
    console.log("MSGGG", this.msg)
    console.log("ALL MSGGG", this.messages)
    console.log("identificaotre", this.identifier)
  },

  methods: {
    formatDate(dateString) {
      const options = { year: 'numeric', month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' };
      const date = new Date(dateString);
      return date.toLocaleDateString('en-US', options);
    },
    async replyToMessage() {
      if (this.text.trim() !== "") {
        try {
          await this.$axios.post("/chats/" + this.msg.chat_id + "/messages", {
            chat_id: this.msg.chat_id ,
            content: this.text,
            forwarded: 0,
            reply: this.msg.message_id,
          });

          this.text = ""
          this.$emit("sentMessage")
        }
        catch (error) {
          console.log(error)
        }
      }
    },

    async deleteMessage() {
        try {
          await this.$axios.delete("/chats/" + this.msg.chat_id + "/messages"+ this.msg.message_id);
          
        }
        catch (error) {
          console.log(error)
        }
      
    },
    async forwardMessage() {
      // this.$emit('forward', this.msg);
      try {
        await this.$axios.post(("/chats/" + this.forwadedToChat + "/forwarded/" + this.msg.messagge_id), {
          chat_name: this.editableChatName
        })
        this.selectedChat.Chat_name = this.editableChatName;

      } catch (error) {
        console.log(error)
      }
    },
    addReaction(reaction) {
      if (this.selectedReaction == reaction) {
        //remove the reaction!
        this.remReac(this.msg.message_id, this.selectedReaction.toString())
      }

      if (this.selectedReaction != -1 && this.selectedReaction != reaction) {
        //remove original and put new reaction
        this.remReac(this.msg.message_id, this.selectedReaction.toString())
        this.addReac(this.msg.message_id, reaction.toString())
      }
    },


    async addReac(msg_id, reaction) {
      try {
        await this.$axios.put(("/chats/" + chat_id + "/messages/" + msg_id + "/reaction/" + reaction))

      } catch (error) {
        console.log(error)
      }
    },

    async remReac(msg_id, reaction) {
      try {
        await this.$axios.delete(("/chats/" + chat_id + "/messages/" + msg_id + "/reaction/" + reaction))

      } catch (error) {
        console.log(error)
      }
    }


  },
};
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

.message-actions {
  margin-top: 10px;
}

.reactions {
  margin-top: 5px;
}

button {
  background-color: transparent;
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
</style>
