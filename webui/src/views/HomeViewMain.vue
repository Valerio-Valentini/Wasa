<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
		}
	},
	methods: {
		
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div class="container-fluid">
        <div class="row">
            <div class="col-4">
                Ingranaggio
            </div>
            <div class="col-4">
                Ricerca
            </div>
            <div class="col-4">
                +
            </div>
        </div>
        <div class="row">

        </div>
	</div>
</template>

<style>
</style>
