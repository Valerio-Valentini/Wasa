<template>
    <div class="col-8">
        <div class="container-fluid" v-if="selectedChat">

            <!-- Chat Header -->
            <div class="row">
                <div class="col-12">
                    <div class="card chat-header d-flex justify-content-between align-items-center p-3">
                        <h5 class="mb-0">{{ getChatName2(selectedChat.Chat_name.split("-")) }}</h5>
                        <button class="btn btn-light btn-sm" @click="openInfo">Info</button>
                    </div>
                </div>
            </div>

            <!-- Chat Messages -->
            <div class="row">
                <div class="col-12">
                    <div class="card chat-box">
                        <div class="card-body chat-messages">
                            <p v-if="!messages || messages.length === 0" class="text-muted">No messages yet...</p>
                            <Message v-for="(msg, _) in messages" :key="msg.message_id" :msg="msg" :messages="messages"
                                :identifier="this.identifier" />
                            <!--@reply="handleReply" @forward="handleForward" @reaction-added="handleReactionAdded"-->
                        </div>
                    </div>
                </div>
            </div>

            <!-- Input Field -->
            <div class="row">
                <div class="col-12">
                    <div class="input-group chat-input">
                        <input type="text" v-model="newMessage" class="form-control" placeholder="Type a message..."
                            @keyup.enter="sendMessage" />
                        <span class="input-group-text send-btn" @click="sendMessage">Send</span>
                    </div>
                </div>
            </div>

            <!-- Info Modal -->
            <div v-if="showInfoModal" class="overlay">
                <div class="modal-content">
                    <h5 v-if="selectedChat.Chat_group">Edit Chat Info</h5>
                    <input v-model="editableChatName" class="form-control mb-2" placeholder="Chat Name"
                        v-if="selectedChat.Chat_group">

                    <h6 v-if="selectedChat.Chat_group">>Members:</h6>
                    <ul class="list-group mb-2" v-if="selectedChat.Chat_group">
                        <li v-for="(member, index) in selectedChat.members" :key="index" class="list-group-item">
                            {{ member }}
                        </li>
                    </ul>

                    <div class="input-group mb-2" v-if="selectedChat.Chat_group">
                        <input v-model="newMember" type="text" class="form-control" placeholder="Add a member..." />
                        <button class="btn btn-primary" @click="addMember">+</button>
                    </div>

                    <div class="input-group mb-2">
                        <button class="btn btn-danger" @click="leaveChatHandler">Leave chat</button>
                    </div>

                    <div class="d-flex justify-content-between">
                        <button class="btn btn-danger" @click="showInfoModal = false">Close</button>
                        <button class="btn btn-success" @click="saveChatName"
                            v-if="selectedChat.Chat_group">Save</button>
                    </div>
                </div>
            </div>

        </div>
    </div>
</template>

<script>
export default {
    props: ["selectedChat", "identifier", "messages"],

    data() {
        return {
            newMessage: "",
            showInfoModal: false,
            editableChatName: "",
            newMember: ""
        };
    },
    methods: {

        openInfo() {
            this.showInfoModal = true
        },

        getChatName2(vector) {
            if (this.identifier == vector[0]) return vector[1]
            return vector[0]
        },

        async sendMessage() {
            if (this.newMessage.trim() !== "") {
                try {
                    let response = await this.$axios.post("/chats/" + this.chat_id + "/messages", {
                        chat_id: this.selectedChat.first_chat_id,
                        content: this.newMessage,
                        forwarded: 0,
                        reply: 0,
                    });

                    this.newMessage = ""
                    this.$emit("sentMessage")
                }
                catch (error) {
                    console.log(error)
                }
            }
        },

        async loadMessage() {
            try {
                let response = await this.$axios.get("/users/" + this.identifier + "/chats/" + this.selectedChat.first_chat_id, {
                });

                this.messages = response.data.messages
            }
            catch (error) {
                console.log(error)
            }
        },


        async addMember() {
            if (this.newMember.trim() !== "") {
                try {
                    await this.$axios.put(("/chats/" + chat_id + "/members/" + this.newMember))
                    this.selectedChat.members.push(this.newMember);
                    this.newMember = "";

                } catch (error) {
                    console.log(error)
                }
            }
        },
        async saveChatName() {
            if (this.editableChatName.trim() !== "") {
                try {
                    await this.$axios.put(("/chats/" + chat_id),{
                        chat_name: this.editableChatName
                    })
                    this.selectedChat.Chat_name = this.editableChatName;

                } catch (error) {
                    console.log(error)
                }
            }
            this.showInfoModal = false;
            this.editableChatName = "";
        },

        async leaveChatHandler() {
            try {
                await this.$axios.delete(("/chats/" + this.selectedChat.first_chat_id + "/members/" + this.identifier))
                alert("You left the chat!")
                this.showInfoModal = false;
            } catch (error) {
                console.log(error)
            }
        }
    },

    async mounted() {
        console.log("INFO SELECT CAHT ", this.selectedChat)
        await this.loadMessage();
    }
};
</script>

<style scoped>
/* Chat Header */
.chat-header {
    background-color: #008069;
    color: white;
    border-radius: 10px;
}

/* Chat Box */
.chat-box {
    height: 400px;
    overflow-y: auto;
    background-color: #f5f5f5;
    padding: 10px;
    border-radius: 10px;
}

/* Messages */
.chat-messages {
    display: flex;
    flex-direction: column;
    gap: 10px;
}

.message {
    background: #dcf8c6;
    padding: 10px;
    border-radius: 10px;
    width: fit-content;
    max-width: 70%;
}

/* Input Field */
.chat-input {
    margin-top: 10px;
}

.send-btn {
    cursor: pointer;
    background-color: #008069;
    color: white;
    border: none;
}

.send-btn:hover {
    background-color: #006a58;
}

/* Overlay Modal */
.overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
}

.modal-content {
    background: white;
    padding: 20px;
    border-radius: 10px;
    width: 300px;
    box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.3);
}
</style>