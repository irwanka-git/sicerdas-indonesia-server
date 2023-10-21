<style type="text/css">
	th.sm{
		font-size: 0.9em!important;
	}

	table.table-x  thead tr  th{
		vertical-align: middle !important;
	}

	table.table-x  tr  td{
		line-height: 1.2em !important;

	}
	table.table-x  tr td p{
		line-height: 1.35em !important;
		margin-top: 0.85rem !important;
	}
</style>
<center>
	<img src="{{url('gambar/'.$data_sesi->avatar)}}" width="48" height="48" class="rounded-circle"><br>
	<b>{{$data_sesi->nama_pengguna}}</b><br>
	<small>{{$data_sesi->organisasi}}</small>

</center>

<div class="row justify-content-center mt-3 mb-2">
	<div class="col-auto">
		<nav class="nav btn-group">
			<a href="#tab1" class="btn btn-outline-primary active" data-bs-toggle="tab">TKD</a>
			<a href="#tab2" class="btn btn-outline-primary" data-bs-toggle="tab">SUASANA KERJA</a>
			<a href="#tab3" class="btn btn-outline-primary" data-bs-toggle="tab">TMI</a>
			<a href="#tab4" class="btn btn-outline-primary" data-bs-toggle="tab">KECERDASAN MAJEMUK</a>
			<a href="#tab5" class="btn btn-outline-primary" data-bs-toggle="tab">MBTI</a>
			<a href="#tab6" class="btn btn-outline-primary" data-bs-toggle="tab">KARAKTERISTIK PRIBADI</a>
		</nav>
	</div>
</div>

<div class="tab-content">
	
	<div class="tab-pane py-2 fade show active" id="tab1">
		<div class="card">
			<div class="card-body">
				@include("report.".$data_sesi->skoring_tabel.".kognitif")
			</div>
		</div>
	</div>

	<div class="tab-pane py-2 fade" id="tab2">
		<div class="card">
			<div class="card-body">
				@include("report.".$data_sesi->skoring_tabel.".suasana-kerja")
			</div>
		</div>
	</div>

	<div class="tab-pane py-2 fade" id="tab3">
		<div class="card">
			<div class="card-body">
				@include("report.".$data_sesi->skoring_tabel.".tmi")
			</div>
		</div>
	</div>

	<div class="tab-pane py-2 fade" id="tab4">
		<div class="card">
			<div class="card-body">
				@include("report.".$data_sesi->skoring_tabel.".kecerdasan-majemuk")
			</div>
		</div>
	</div>

	<div class="tab-pane py-2 fade" id="tab5">
		<div class="card">
			<div class="card-body">
				@include("report.".$data_sesi->skoring_tabel.".mbti")
			</div>
		</div>
	</div>

	<div class="tab-pane py-2 fade" id="tab6">
		<div class="card">
			<div class="card-body">
				@include("report.".$data_sesi->skoring_tabel.".karakter-pribadi")
			</div>
		</div>
	</div>
 

	<div class="tab-pane py-2 fade" id="saran">
		<div class="card">
			<div class="card-body">
				@include("report.".$data_sesi->skoring_tabel.".saran")
			</div>
		</div>
	</div>
</div>