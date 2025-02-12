<template>
    <div class="user-item">
      <span class="user-name">{{ user.user_id }}</span>
      <button class="btn btn-primary" :disabled="checkExistanceChat(user.user_id)" @click="chat">{{label}}</button>
    </div>
  </template>
  
  <script>
  export default {
    props: ["user","chats", "label"],
    methods: {
      chat() {
        this.$emit("chat", this.user);
      },

      checkExistanceChat(element){
        return this.chats.some(chat => chat.Chat_name.split("-")[0]==element|| chat.Chat_name.split("-")[1]==element);
      }
    }
  };
  </script>
  
  <style scoped>
  .user-item {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 10px;
    border: 1px solid #ddd;
    border-radius: 5px;
    margin-bottom: 5px;
    background: #f9f9f9;
  }
  .user-name {
    flex-grow: 1;
    font-weight: bold;
  }
  button {
    margin-left: 10px;
  }
  </style>
  