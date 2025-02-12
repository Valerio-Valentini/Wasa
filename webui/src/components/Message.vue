<template>
    <div :class="['message', messageType]">
      <div class="message-info">
        <span class="user-id me-2">{{ msg.user_id }}</span>
        <span class="message-date">{{ formatDate(msg.date) }}</span>
      </div>
      <div class="message-content">
        {{ msg.content }}
      </div>
      <div class="message-actions">
        <button @click="replyToMessage">Reply</button>
        <button @click="forwardMessage">Forward</button>
      </div>
      <div class="reactions">
        <button v-for="reaction in reactions" :key="reaction" @click="addReaction(reaction)">
          {{ reaction }}
        </button> <!--{{ reaction }} ({{ msg.reactions[reaction] || 0 }})-->
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: ["msg"],
    data() {
      return {
        reactions: ['0', '1', '2', '3'], // Reactions for each message
      };
    },

    computed: {

      messageType() {
        if (this.msg.reply > 0) {
          return 'reply';
        } else if (this.msg.forwarded) {
          return 'forwarded';
        } else {
          return 'regular';
        }
      },
    },
    methods: {
      // Format the date to a more readable format
      formatDate(dateString) {
        const options = { year: 'numeric', month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' };
        const date = new Date(dateString);
        return date.toLocaleDateString('en-US', options);
      },
  
      // Handle reply action
      replyToMessage() {
        this.$emit('reply', this.msg); // Emit reply event to parent
      },
  
      // Handle forward action
      forwardMessage() {
        this.$emit('forward', this.msg); // Emit forward event to parent
      },
  
      // Add a reaction to the message
      addReaction(reaction) {
        if (!this.msg.reactions) this.msg.reactions = {}; // Initialize reactions if not yet set
        if (!this.msg.reactions[reaction]) {
          this.msg.reactions[reaction] = 1;
        } else {
          this.msg.reactions[reaction] += 1;
        }
        this.$emit('reaction-added', { message: this.msg, reaction }); // Emit the reaction change
      },
    },
  };
  </script>
  
  <style scoped>
  /* Message Styles */
  .message {
    padding: 10px;
    border-radius: 10px;
    margin: 10px 0;
  }
  
  .message.regular {
    background-color: #e5e5e5;
  }
  
  .message.forwarded {
    background-color: #dcf8c6;
  }
  
  .message.reply {
    background-color: #f0f0f0;
    margin-left: 20px;
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
  </style>
  