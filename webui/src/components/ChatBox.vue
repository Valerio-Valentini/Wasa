<template>
    <div class="col-8">
        <div class="container-fluid" v-if="selectedChat">

            <!-- Chat Header -->
            <div class="row">
                <div class="col-12">
                    <div class="card chat-header d-flex justify-content-between align-items-center p-3">
                        <h5 class="mb-0">{{ getChatName2(selectedChat.chat_name.split("-")) }} </h5>
                        <button class="btn btn-light btn-sm" @click="openInfo">Info</button>
                        <div>
                        <img  v-if="selectedChat.chat_group" id="photo" :src="source + '/chats/' + selectedChat.first_chat_id + '/photo'">
                        <img v-else id="photo" :src="source + '/users/' + getChatName2(selectedChat.chat_name.split('-')) + '/photo'">
                    </div>
                    </div>
                </div>
            </div>

            <!-- Chat Messages -->
            <div class="row">
                <div class="col-12">
                    <div class="card chat-box">
                        <div v-if="showImagePopUp" class="popup-overlay custom">
                            <div class="popup-content">
                                <div class="input-group flex-nowrap">
                                    <input id="fileUploader" type="file" class="profile-file-upload" accept=".jpg">
                                    <button class="btn btn-primary" @click="uploadFile">
                                         Send Photo
                                    </button>
                                </div>    
                                <button class="btn btn-danger" @click="showImagePopUp = false">Close</button>
                            </div>
                        </div>
                        <div class="card-body chat-messages">
                            <p v-if="!messages || messages.length === 0" class="text-muted">No messages yet...</p>
                            <div v-else class="row" v-for="(msg, i) in messages" :key="i"> 
                                <Message v-if="msg.photo_id == -1" :key="msg.message_id" :msg="msg" :messages="messages"
                                :identifier="this.identifier" :chats="this.chats" @newMessage="sendNewChat"/>
                                <MediaMessage v-else :key="msg.message_id +1" :msg="msg" :messages="messages"
                                :identifier="this.identifier" :chats="this.chats" @newMessage="sendNewChat"/>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Input Field -->
            <div class="row">
                <div class="col-12">
                    <div class="input-group chat-input">
                        <button @click="showImagePopUp = true" class="btn btn-primary">+</button>
                        <input type="text" v-model="newMessage" class="form-control" placeholder="Type a message..."
                            @keyup.enter="sendMessage" />
                        <span class="input-group-text send-btn" @click="sendMessage">Send</span>
                    </div>
                </div>
            </div>

            <!-- Info Modal -->
            <div v-if="showInfoModal" class="overlay">
                <div class="modal-content">
                    <h5 v-if="this.selectedChat.chat_group">Edit Chat Info</h5>
                    <input v-model="editableChatName" class="form-control mb-2" placeholder="Chat Name"
                        v-if="this.selectedChat.chat_group">

                    <h6 v-if="this.selectedChat.chat_group">Members:</h6>
                    
                    
                    <ul class="list-group mb-2" v-if="this.selectedChat.chat_group">
                        <li v-for="(member, index) in this.selectedChat.chat_members.split('-')" :key="index" class="list-group-item">
                            {{ member }}
                        </li>
                    </ul>
                    

                    <h6 v-if="this.selectedChat.chat_group">Add Members:</h6>
                    <div class="input-group mb-2" v-if="this.selectedChat.chat_group">
                        <input v-model="newMember" type="text" class="form-control" placeholder="Add a member..." />
                        <button class="btn btn-primary" @click="addMember">+</button>
                    </div>

                    <div class="input-group mb-2">
                        <button class="btn btn-danger" @click="leaveChatHandler">Leave chat</button>
                    </div>

                    <div class="d-flex justify-content-between">
                        <button class="btn btn-danger" @click="showInfoModal = false">Close</button>
                        <button class="btn btn-success" @click="saveChatName"
                            v-if="this.selectedChat.chat_group">Save</button>
                    </div>
                </div>
            </div>

        </div>
    </div>
</template>

<script>
export default {
    props: ["selectedChat", "identifier", "messages", "chats"],

    data() {
        return {
            newMessage: "",
            showInfoModal: false,
            editableChatName: "",
            newMember: "",
            showImagePopUp: false,
            source: __API_URL__

        };
    },
    methods: {

        openInfo() {
            this.showInfoModal = true
            console.log(this.selectedChat)
        },

        getChatName2(vector) {
            if (vector.length == 1) return vector[0]
            if (this.identifier == vector[0]) return vector[1]
            return vector[0]
        },

        async sendMessage() {
            if (this.newMessage.trim() !== "") {
                try {
                    await this.$axios.post("/chats/" + this.chat_id + "/messages", {
                        chat_id: this.selectedChat.first_chat_id,
                        content: this.newMessage,
                        forwarded: 0,
                        reply: 0,
                        photo_id: -1,
                    });

                    this.newMessage = ""
                    this.$emit("newMsg")
                    this.$emit("sentMessage")
                }
                catch (error) {
                    console.log(error)
                }
            }
        },


        async addMember() {
            if (this.newMember.trim() !== "") {
                try {
                    await this.$axios.put(("/chats/" + this.selectedChat.first_chat_id + "/members/" + this.newMember))
                    alert("You added " + this.newMember)
                    this.newMember = "";

                } catch (error) {
                    if (error.response.status == 404) {
                        alert(this.newMember + " doesn't exist!")
                        this.newMember = "";
                    } else {
                        console.log(error)
                    }
                }
            }
        },
        async saveChatName() {
            if (this.editableChatName.trim() !== "") {
                try {
                    await this.$axios.put(("/chats/" + this.selectedChat.first_chat_id), {
                        chat_name: this.editableChatName
                    })
                    // this.selectedChat.chat_name = this.editableChatName;

                    //localStorage.setItem("selectedChat", JSON.stringify(this.selectedChat))

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
                localStorage.removeItem("selectedChat")
                window.location.reload(); 
            } catch (error) {
                console.log(error)
            }
        },

        sendNewChat() {
            this.$emit("newMessage")
        },

        async uploadFile() {
            let fileInput = document.getElementById("fileUploader")
            console.log(fileInput)
            const file = fileInput.files[0]
            const reader = new FileReader()
            reader.readAsArrayBuffer(file)
            // /chats/:chat_id/media
            reader.onload = async () => {
                let response = await this.$axios.post("/chats/" + this.selectedChat.first_chat_id + "/media", reader.result, {
                    headers: {
                        "Content-Type": file.type
                    }
                })
                let photo_id = response.data.photo_id
                await this.$axios.post("/chats/" + this.selectedChat.first_chat_id + "/messages", {
                        chat_id: this.selectedChat.first_chat_id,
                        content: this.newMessage,
                        forwarded: 0,
                        reply: 0,
                        photo_id: parseInt(photo_id, 10),
                    });
                    this.$emit("newMsg")
                    this.$emit("sentMessage")

            }
        }
    },
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

.custom {
    position: absolute;
    top: 50%;
    left: 50%;
    background-color: gray;
}
</style>