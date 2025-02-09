<template>
    <div class="user-item">
      <span class="user-name">{{ user.user_id }}</span>
      <button class="btn btn-primary" @click="addToGroup">Add to Group</button>
      <button class="btn btn-secondary" @click="chat">Chat</button>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      user: {
        type: Object,
        required: true
      }
    },
    methods: {
      addToGroup() {
        this.$emit("add-to-group", this.user);
      },
      chat() {
        this.$emit("chat", this.user);
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
  