<template>
    <div class="contaienr-fluid">
        <div class="row">
            <nav class="navbar navbar-expand-lg bg-body-tertiary">
                <div class="container-fluid">
                    <a class="navbar-brand" href="#">WasaText</a>
                    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
                        <span class="navbar-toggler-icon"></span>
                    </button>
                    <div class="collapse navbar-collapse" id="navbarNavDropdown">
                        <ul class="navbar-nav">
                            <li class="nav-item">
                                <a class="nav-link" aria-current="page" href="#">Profile</a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link" href="#">LogOut</a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link" href="#">NewGroup</a>
                            </li>
                            <li class="nav-item dropdown">
                                <nav class="navbar bg-body-tertiary">
                                    <div class="container-fluid">
                                        <form class="d-flex" role="search">
                                            <input class="form-control me-2" type="search" placeholder="@username" aria-label="Search">
                                            <button class="btn btn-outline-success" type="submit">Search</button>
                                        </form>
                                    </div>
                                </nav>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
        </div>
        <div class="row">
            <div class="col-4">
                <div class="list-group">
                    <a href="#" class="list-group-item list-group-item-action active" aria-current="true">
                        The current link item
                    </a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                    <a href="#" class="list-group-item list-group-item-action">link item</a>
                </div>
            </div>
            <div class="col-8">
                <div class="container-fluid">
                    <div class="row">
                        <div class="col-12">
                            <div class="card" id="spc">
                                
                                <div class="card-body">
                                    <p class="card-text">testo</p>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col-12">
                            <div class="input-group" id="txt">
                                <input type="text" class="form-control" aria-label="Dollar amount (with dot and two decimal places)">
                                <span class="input-group-text" id="btn">Invia</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
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