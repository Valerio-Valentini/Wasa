<script>
export default {
    data: function () {
        return {
            username: null,
            chats: [],
            searchUser: null,
            foundUsers: [],
            chatSelected: null,
            messages: []
        }
    },

    methods: {

        logout() {
            this.$emit("logout");
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
                let response = await this.$axios.put("/chats", {
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

        async update_chats(){
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

        async selectedChatHandler(pickedChat){
            try {
                this.chatSelected = pickedChat;
                let response = await this.$axios.get("/users/" + this.identifier + "/chats/" + this.chatSelected.first_chat_id, {
                });

                this.messages = response.data.messages
                
            }
            catch (error) {
                console.log(error)
            }
        },
        
    },

    async mounted() {
        this.update_chats();
    },

    props: ["identifier"]
}
</script>

<template>
    <div class="contaienr-fluid">
        <div class="row">
            <nav class="navbar navbar-expand-lg bg-body-tertiary">
                <div class="container-fluid">
                    <a class="navbar-brand rounded text-center" style="color: white; background-color: #008069; width: 140px">WasaText</a>
                    <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
                        data-bs-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false"
                        aria-label="Toggle navigation">
                        <span class="navbar-toggler-icon"></span>
                    </button>
                    <div class="collapse navbar-collapse" id="navbarNavDropdown">
                        <ul class="navbar-nav">
                            <li class="nav-item">
                                <a class="nav-link" aria-current="page">{{ identifier }}</a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link hover-box" @click="logout">LogOut</a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link" href="#">NewGroup</a>
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
                <h5 class="offcanvas-title" id="offcanvasWithBothOptionsLabel" v-if="this.foundUsers.length != 0">Utenti
                    trovati:</h5>
                <h5 class="offcanvas-title" id="offcanvasWithBothOptionsLabel" v-else>Nessun utente trovato!</h5>
                <button type="button" class="btn-close" data-bs-dismiss="offcanvas" aria-label="Close"></button>
            </div>
            <div class="offcanvas-body">
                <UserItem v-for="(user, index) in foundUsers" :key="index" :user="user"
                    @chat="handleChat(user)" />

            </div>
        </div>


        <div class="row">
            <div class="col-4">
                <div class="list-group">
                    <a class="list-group-item list-group-item-action hover-box" v-for="(chat, index) in chats"
                        :key="index" @click="selectedChatHandler(chat)" >{{ chat.Chat_name.split("-")[1]  }} </a>
                </div>
            </div>
            <ChatBox :selectedChat="this.chatSelected" :identifier="this.identifier" :messages="this.messages"/>
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