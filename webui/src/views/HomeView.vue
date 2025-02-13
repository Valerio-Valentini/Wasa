<script>
export default {
    data: function () {
        return {
            username: null,
            chats: [],
            searchUser: null,
            foundUsers: [],
            chatSelected: null,
            messages: [],
            lastMsgPerChat: []
        }
    },

    methods: {

        logout() {
            this.$emit("logout");
            localStorage.removeItem("selectedChat")
        },

        goToCreateGroup() {
            this.$router.push("/createGroup")
        },

        async search() {
            try {
                let response = await this.$axios.get("/users", {
                    params: {
                        user_id: this.searchUser.trim()
                    }
                });
                this.foundUsers = response.data
                this.searchUser = null
            }
            catch (error) {
                console.log(error)
            }
        },

        async handleChat(user) {
            try {
                await this.$axios.put("/chats", {
                    user_ids: [this.identifier.toString(), user.user_id.toString()],
                    chat_group: false
                }
                );

                this.update_chats();
            }
            catch (error) {
                console.log(error)
            }

        },

        async update_chats() {
            try {
                let response = await this.$axios.get("/users/" + this.identifier + "/chats", {
                    user_id: this.identifier,
                    photo_id: 0
                },
                )
                this.chats = response.data.chats
            }
            catch (error) {
                console.log(error)
            }
        },

        async selectedChatHandler(pickedChat) {
            try {
                if (pickedChat) {
                    this.chatSelected = pickedChat;
                    localStorage.setItem("selectedChat", JSON.stringify(pickedChat))
                }
                let response = await this.$axios.get("/users/" + this.identifier + "/chats/" + this.chatSelected.first_chat_id, {
                });

                this.messages = response.data.messages

            }
            catch (error) {
                console.log(error)
            }
        },

        getChatName(vector) {
            if (this.identifier === vector[0]) return vector[1]
            return vector[0]
        },

        async getLastMsg(chat_id) {
            try {
                if (chat_id) {

                    let response = await this.$axios.get("/users/" + this.identifier + "/chats/" + this.chatSelected.first_chat_id, {
                    });

                    let arr = response.data.messages
                    if (arr.length > 0) arr = arr[arr.length - 1]
                    return arr
                }
            }
            catch (error) {
                console.log(error)
            }
        },

        async loadLastMessages() {
            this.lastMsgPerChat = [];
            for (var chat in this.chats) {
                let x = await this.getLastMsg(this.chats[chat].first_chat_id);
                if (!x) continue
                this.lastMsgPerChat.push(x)
            }
        },

        truncateString(str, maxLength = 10) {
            return str.length > maxLength ? str.slice(0, maxLength) + "..." : str;
        },
        formatDate(isoString) {
            const date = new Date(isoString);

            // Check if the date is valid by checking if getTime() is NaN
            if (isNaN(date.getTime())) {
                return ''; // Return empty string if invalid date
            }

            return date.toLocaleString('en-US', {
                year: 'numeric',
                month: 'long',
                day: 'numeric',
                hour: '2-digit',
                minute: '2-digit',
                second: '2-digit',
                hour12: true
            });
        },

        util1(chat) {
            let messageContent = " ";
            for (let i = 0; i < this.lastMsgPerChat.length; i++) {
                if (this.lastMsgPerChat[i].chat_id === chat.first_chat_id) {
                    messageContent = this.lastMsgPerChat[i].content;
                    return messageContent;
                }
            }
            return messageContent;
        }

    },

    async mounted() {
        await this.update_chats();
        await this.loadLastMessages()
        this.selectedChatHandler(JSON.parse(localStorage.getItem("selectedChat")))



    },

    props: ["identifier"]
}
</script>

<template>
    <div class="contaienr-fluid">
        <div class="row">
            <nav class="navbar navbar-expand-lg bg-body-tertiary">
                <div class="container-fluid">
                    <a class="navbar-brand rounded text-center"
                        style="color: white; background-color: #008069; width: 140px">WasaText</a>
                    <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
                        data-bs-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false"
                        aria-label="Toggle navigation">
                        <span class="navbar-toggler-icon"></span>
                    </button>
                    <div class="collapse navbar-collapse" id="navbarNavDropdown">
                        <ul class="navbar-nav">
                            <li class="nav-item">
                                <a class="nav-link" style="font-style: italic; font-weight: bold;"
                                    aria-current="page">{{ identifier }}</a>
                            </li>
                            <li class="nav-item me-2">
                                <button class="btn btn-danger" @click="logout">LogOut</button>
                            </li>
                            <li class="nav-item">
                                <button class="btn btn-warning" @click="goToCreateGroup">Create Group</button>
                            </li>
                            <li class="nav-item dropdown">
                                <nav class="navbar bg-body-tertiary">
                                    <div class="container-fluid">
                                        <form class="d-flex" role="search">
                                            <input class="form-control me-2" type="search" placeholder="@username"
                                                aria-label="Search" v-model="searchUser">
                                            <button class="btn btn-outline-success" type="submit"
                                                data-bs-toggle="offcanvas" data-bs-target="#offcanvasWithBothOptions"
                                                aria-controls="offcanvasWithBothOptions"
                                                :disabled="!this.searchUser || this.searchUser.trim().length == 0"
                                                @click="search">Search</button>
                                        </form>
                                    </div>
                                </nav>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
        </div>


        <div class="offcanvas offcanvas-start" data-bs-scroll="true" tabindex="-1" id="offcanvasWithBothOptions"
            aria-labelledby="offcanvasWithBothOptionsLabel">
            <div class="offcanvas-header">
                <h5 class="offcanvas-title" id="offcanvasWithBothOptionsLabel"
                    v-if="this.foundUsers && this.foundUsers.length != 0">Utenti
                    trovati:</h5>
                <h5 class="offcanvas-title" id="offcanvasWithBothOptionsLabel" v-else>Nessun utente trovato!</h5>
                <button type="button" class="btn-close" data-bs-dismiss="offcanvas" aria-label="Close"></button>
            </div>
            <div class="offcanvas-body">
                <UserItem v-for="(user, index) in foundUsers" :key="index" :user="user" @chat="handleChat(user)"
                    :chats="this.chats" :label="'Chat'" />

            </div>
        </div>


        <div class="row">
            <div class="col-4">
                <div class="list-group">
                    <a class="list-group-item list-group-item-action hover-box" v-for="(chat, index) in chats"
                        :key="index" @click="selectedChatHandler(chat)">

                        <!--{{ util1(chat) }}-->    
                        <div class="row" style="font-weight: bolder;">
                            {{ getChatName(chat.Chat_name.split("-")) }}
                        </div>
                        <div class="row" style="font-style: italic;">
                         >> {{ lastMsgPerChat.find(msg => msg.chat_id === chat.first_chat_id)?.content ?? " " }}
                        </div>
                        <div class="row" style="font-style: italic;">
                            {{ formatDate(lastMsgPerChat.find(msg => msg.chat_id === chat.first_chat_id)?.date) }}
                        </div>
                    </a>
                </div>
            </div>
            <ChatBox @sentMessage="selectedChatHandler" :selectedChat="this.chatSelected" :identifier="this.identifier"
                :messages="this.messages" :chats="this.chats" @newMsg="loadLastMessages" />
        </div>
    </div>
</template>

<style>
#spc {
    height: 730px;
    margin-top: 10px;
    margin-left: -25px;
}

#txt {
    margin-left: -25px;
    margin-top: 5px;
    height: 75px;
    width: 865px;
}

#btn {
    background-color: rgb(13, 110, 253);
    color: white;
}
</style>