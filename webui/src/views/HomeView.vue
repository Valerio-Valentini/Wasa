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
            <div class="col-4 d-flex justify-content-start">
                Ingranaggio
            </div>
            <div class="col-4 d-flex justify-content-center">
               <!-- As a link -->
				<nav class="navbar bg-body-tertiary">
  					<div class="container-fluid">
    					<a class="navbar-brand" href="#">Ricerca</a>
  					</div>
				</nav>
            </div>
            <div class="col-4 d-flex justify-content-end">
                +
            </div>
        </div>
        <div class="row">
			<div class="col-4">
				<div class="container-fluid">
					<div class="row">
						<div class="card">
  							<div class="card-body">
    							Chat1
  							</div>
						</div>
					</div>
					<div class="row">
						Chat2
					</div>
					<div class="row">
						Chat3
					</div>
					<div class="row">
						Chat4
					</div>
				</div>
			</div>
			<div class="col-8">
				<div class="container-fluid">
					Contenuto chat x
				</div>
			</div>
        </div>
	</div>
</template>

<style>
</style>