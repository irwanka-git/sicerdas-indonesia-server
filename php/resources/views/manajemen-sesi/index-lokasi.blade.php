<?php
// tabel: quiz_sesi_template
// quiz_sesi_template
// id_quiz_template int
// nama_sesi        varchar
// skoring_tabel    varchar
// uuid             char


loadHelper('akses,function');

$id_role_biro = 5;
$main_path = Request::segment(1);
$isAdminSistem = isAdminSistem();
$isAdminBiro = isAdminBiro();
?>
@extends('layout')
@section("css")
<link rel="stylesheet" href="//cdn.quilljs.com/1.3.6/quill.snow.css" />
@endsection
@section("pagetitle")

@endsection

@section('content')

 <div class="container-fluid p-0">
 					@if($isAdminSistem)
					<a href="#" data-bs-toggle="modal" data-bs-target="#modal-tambah" class="btn btn-primary float-end mt-n1"><i class="fas fa-plus"></i> Sesi Tes Baru</a>
					@endif
					<h1 class="h3 mb-3">Manajemen Sesi Tes</h1>
					<nav style="--bs-breadcrumb-divider: '>';" aria-label="breadcrumb">
					  <ol class="breadcrumb">
					    <li class="breadcrumb-item"><a href="{{url('/manajemen-sesi')}}">Manajemen Sesi Tes</a></li>
					    <li class="breadcrumb-item"><a href="{{url('/manajemen-sesi/explore/'.$data_tahun['tahun'])}}">Tahun {{$data_tahun['tahun']}}</a></li>
					    <li class="breadcrumb-item"><a href="{{url('/manajemen-sesi/explore/'.$data_tahun['tahun'].'/'.$data_biro['id_user_biro'])}}">{{$data_biro['nama_biro']}} ({{$data_biro['jumlah_tes']}})</a></li>
					    <li class="breadcrumb-item active"> {{$data_provinsi->name}} ({{$data_provinsi->jumlah_tes}})</li>
					  </ol>
					</nav>
					<div class="row">
						<div class="col-12" style="margin-top: 20px;">

							<table id="tabel-lokasi" class="table table-striped table-hover table-sm" style="width:100%">
								<thead>
									<tr>							
										<th style="text-align: center;"  align="center" width="5%">No</th>
										<th width="30%">Lokasi</th>
										<th width="30%">Kabupaten / Kota</th>
										<th style="text-align: center;"  width="15%">Jumlah Sesi Tes</th>
										<th style="text-align: center;"  width="15%">Tanggal</th>
									</tr>
								</thead>
								<tbody>
									<?php $no=1;?>
									@foreach($list_lokasi as $r)
										<?php 
										//$url_detil = url('/manajemen-sesi/explore/'.$r->tahun);
										$url_detil = url('/manajemen-sesi/explore/'.$data_tahun['tahun'].'/'.$data_biro['id_user_biro'].'/'.$data_provinsi->id.'/'.$r->id_lokasi);
										?>
										<tr>
											<td align="center">{{$no++}}</td>
											<td><a href="{{$url_detil}}">{{$r->nama_lokasi}}</a></td>
											<td>{{$r->kabupaten}}</td>
											<td align="center">{{$r->jumlah_tes}}</td>
											<td align="center">{{$r->tanggal!="" ? tgl_indo($r->tanggal) : "Belum Diatur"}}</td>
										</tr>
									@endforeach
								</tbody>
							</table>

						</div>
						
						@if(count($list_lokasi)==0)
						<div class="col-12">
							<div class="alert alert-primary mb-5" role="alert">
								<div class="alert-message">
									<strong>Perhatian:</strong> Untuk saaat sesi tes belum tersedia, silahkan hubungi admin <b>Si Cerdas Indonesia</b> untuk mendaftarkan sesi tes yang akan diselenggarakan!.
								</div>
							</div>
						</div>
						@endif
					</div>
				</div>

@endsection

@section("js")
<script src="//cdn.quilljs.com/1.3.6/quill.min.js"></script>
<script type="text/javascript">
	$(function() {
		//var toolbarOptions = [['bold', 'italic'],['link', 'image']];
		$('#tabel-lokasi').DataTable({"iDisplayLength": 10,});
	})
</script>
@if($isAdminSistem)
@include("manajemen-sesi.form-tambah")
@endif
@endsection