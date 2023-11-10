<?php
// tabel: quiz_sesi_template
// quiz_sesi_template
// id_quiz_template int
// nama_sesi        varchar
// skoring_tabel    varchar
// uuid             char

$main_path = Request::segment(1);
loadHelper('akses,function');
$list_mode = get_list_enum_values('quiz_sesi_master', 'mode');
$list_jenis = get_list_enum_values('quiz_sesi','jenis');
// dd($list_mode);
?>
@extends('layout')
@section("css")
<link rel="stylesheet" href="//cdn.quilljs.com/1.3.6/quill.snow.css" />
@endsection
@section("pagetitle")

@endsection

@section('content')

<div class="row">
	<div class="col-12">
		<div class="card">
			<div class="card-header">
				<h5 class="card-title">Jenis Tes</h5>
				<h6 class="card-subtitle text-muted">Fitur Ini Digunakan Untuk Manajemen Data Master Jenis Tes Si-Cerdas </h6>
			</div>
			<div class="card-body">
				@if(ucc())
				{{Html::btnModal('<i class="la la-plus-circle"></i> Tambah Jenis','modal-tambah','primary')}}
				<hr>
				@endif
				<table id="datatable" class="table table-striped table-hover table-sm" style="width:100%">
					<thead>
						<tr>							
							<th width="10%">Kode</th>
							<th width="35%">Nama Jenis Tes</th>
							<th width="10%">Jenis</th>
							<th width="10%">Gambar</th>
							<th width="10%">Quiz/Tes</th>
							<th width="15%">Sesi</th>
							<th width="15%">Format</th>							
							@if(ucu() || ucd())
							<th width="10%">Actions</th>
							@endif
						</tr>
					</thead>
					<tbody>

					</tbody>
				</table>
			</div>
		</div>
	</div>
</div>

@endsection

@section("modal")


@if(ucc())
<!-- MODAL FORM TAMBAH -->
{{ Form::bsOpen('form-tambah',url($main_path."/insert")) }}
	{{Html::mOpenLG('modal-tambah','Tambah Jenis Tes')}}
		{{ Form::bsTextField('Kode Tes','kode','',true,'md-8') }}
		{{ Form::bsTextField('Nama Tes','nama_sesi','',true,'md-8') }}
		{{ Form::bsSelect2('Jenis','jenis',$list_jenis,'',true,'md-8')}}
		<div class="mb-3">
			<small>Ukuran gambar yang direkomendasikan adalah 400 x 250 pixel</small>
			<div class="input-group">
				<button id="btn_gambar" class="btn btn-primary btn-upload-gambar" data-field="gambar" 
					data-form="form-tambah" type="button"><i class="la la-upload"></i> Upload Gambar</button>
				<input type="text" required="required"  
					name="gambar" id="gambar" class="form-control">
				<button data-field="gambar" data-form="form-tambah" class="btn btn-secondary btn-lihat-gambar" type="button"><i class="la la-eye"></i> Lihat</button>
			</div>
		</div>
		<div class="mb-3">
			<label class="form-label">Pendahuluan <star>*</star> </label>
			<div  id="pendahuluan_tambah"></div>
			<textarea style="display: none;" name="pendahuluan" id="pendahuluan"  required="required"></textarea>
		</div>
	{{Html::mCloseSubmitLG('Simpan')}}
{{ Form::bsClose()}}
@endif

{{Html::mOpenLG('modal-using','Quiz/Tes yang menggunakan jenis tes ini.')}}
	<div id="panel-using"   height="500px" width="100%" title=""></div>
{{Html::mCloseLG()}}

@if(ucu())
<!-- MODAL FORM EDIT -->
{{ Form::bsOpen('form-edit',url($main_path."/update")) }}
	{{Html::mOpenLG('modal-edit','Edit Jenis Tes')}}
		{{ Form::bsTextField('Kode Tes','kode','',true,'md-8') }}
		{{ Form::bsTextField('Nama Sesi','nama_sesi','',true,'md-8') }}
		{{ Form::bsSelect2('Jenis','jenis',$list_jenis,'',true,'md-8')}}
		<div class="mb-3">
			<small>Ukuran gambar yang direkomendasikan adalah 400 x 250 pixel</small>
			<div class="input-group">
				<button id="btn_gambar" class="btn btn-primary btn-upload-gambar" data-field="gambar" 
					data-form="form-edit" type="button"><i class="la la-upload"></i> Upload Gambar</button>
				<input type="text" required="required"  
					name="gambar" id="gambar" class="form-control">
				<button data-field="gambar" data-form="form-edit" class="btn btn-secondary btn-lihat-gambar" type="button"><i class="la la-eye"></i> Lihat</button>
			</div>
		</div>
		<div class="mb-3">
			<label class="form-label">Pendahuluan <star>*</star> </label>
			<div  id="pendahuluan_edit"></div>
			<textarea style="display: none;" name="pendahuluan" id="pendahuluan"  required="required"></textarea>
		</div>
		{{ Form::bsHidden('uuid','') }}
	{{Html::mCloseSubmitLG('Simpan')}}
{{ Form::bsClose()}}
@endif


@if(ucd())
<!-- FORM DELETE -->
{{ Form::bsOpen('form-delete',url($main_path."/delete")) }}
	{{ Form::bsHidden('uuid','') }}
{{ Form::bsClose()}}
@endif


{{ Form::bsOpen('form-upload-gambar',url($main_path."/upload-gambar")) }}
	 <input type="file" style="display: none;" id="upload-gambar" name="image" accept="image/*">
{{ Form::bsClose()}}


@endsection

@section("js")
<script src="//cdn.quilljs.com/1.3.6/quill.min.js"></script>
<script type="text/javascript">
	$(function() {
		$(".select2").selectize();
		var toolbarOptions = [['bold', 'italic']];
		var $tabel1 = $('#datatable').DataTable({
			processing: true,
			responsive: true,
			fixedHeader: true,
			serverSide: true,
			ajax: "{{url($main_path.'/dt')}}",
			"iDisplayLength": 25,
			columns: [
				{data: 'kode',name: "kode",orderable: false,searchable: false,sClass: "text-center"},
				{data: 'nama_sesi',name: "nama_sesi",orderable: false,searchable: false,sClass: ""},
				{data: 'jenis',name: "jenis",orderable: false,searchable: false,sClass: "text-center"},
				{data: 'gambar',name: "gambar",orderable: false,searchable: false,sClass: ""},
				{data: 'quiz_using',name: "quiz_using",orderable: false,searchable: false,sClass: ""},
				{data: 'jumlah_sesi',name: "jumlah_sesi",orderable: false,searchable: false,sClass: "text-center"},
				{data: 'item_report',name: "item_report",orderable: false,searchable: false,sClass: "text-center"},
				@if(ucu() || ucd())
				{data: 'action',orderable: false,searchable: false,sClass: "text-center"},
				@endif
			],
				"fnRowCallback": function(nRow, aData, iDisplayIndex, iDisplayIndexFull) {
					$(nRow).addClass(aData["rowClass"]);
					return nRow;
			},
			"drawCallback": function(settings) {
				initTooltips();
				@if(ucd())
				initKonfirmDelete();
				@endif
				initViewGambar();
				initViewUsing();
			}
		});

		$tabel1.on('responsive-display', function(e, datatable, columns) {
			initKonfirmDelete();;
		});

		
		//HANDLE UPLOAD GAMBAR
		 @if(ucc() || ucu() || ucd())
			 $field_gambar = '';
			 $form_gambar = '';
			 $field_gambar_lihat = '';
			 $form_gambar_lihat = '';
			 $base_url_image = '{{url("gambar")}}';

			 $(".btn-upload-gambar").on('click', function(){
			 	$field_gambar = $(this).data('field');
			 	$form_gambar = $(this).data('form');
			 	$("#upload-gambar").trigger('click');
			 });

			  $(".btn-lihat-gambar").on('click', function(){
				 	$field_gambar_lihat = $(this).data('field');
				 	$form_gambar_lihat = $(this).data('form');
				 	$filename = $("#"+$form_gambar_lihat +" #"+$field_gambar_lihat).val();
				 	if($filename){
			 			$.alert({
						    title: 'Gambar',
						    columnClass: 'col-md-6',
						    content: '<div width="100%"><center><img  width="100%" src="'+$filename+'" class="img-fluid rounded-lg" ></center></div>',
						});
				 	}
			 });

		 	$("#upload-gambar").on('change', function(){
		 		if($(this).val()){
		 			$("#form-upload-gambar").submit();
		 		}
			});

		 	$('#form-upload-gambar').ajaxForm({
				beforeSubmit:function(){ disableButton("#"+$form_gambar +" #btn_"+$field_gambar) },
				success:function($respon){
					enableButton("#"+$form_gambar +" #btn_"+$field_gambar);
					if ($respon.status==true){
						  $("#upload-gambar").val('');
						  $("#"+$form_gambar +" #"+$field_gambar).val($respon.data);
					}else{
						errorNotify($respon.message);
					}
				},
				error:function(){errorNotify('Terjadi Kesalahan!');}
			}); 
		 @endif

		@if(ucc())

		var quill_pendahuluan_tambah = new Quill('#pendahuluan_tambah', {
			placeholder: 'Tulis Salam Pembuka / Pendahuluan.....',
			theme: 'snow',
			modules: {
				toolbar: toolbarOptions
			}
		});
		quill_pendahuluan_tambah.on('text-change', function(delta, oldDelta, source) {
			$("#form-tambah #pendahuluan").val($("#pendahuluan_tambah .ql-editor").html());
		});

		

		$validator_form_tambah = $("#form-tambah").validate();
		$("#modal-tambah").on('show.bs.modal', function(e) {
			$validator_form_tambah.resetForm();
			$("#form-tambah").clearForm();
			//$("#petunjuk_tambah .ql-editor").html('');
			$("#form-tambah #petunjuk_sesi").val('');
			$('#form-tambah #jenis').selectize()[0].selectize.clear();
			enableButton("#form-tambah button[type=submit]")
		});

		$('#form-tambah').ajaxForm({
			beforeSubmit: function() {
				disableButton("#form-tambah button[type=submit]")
			},
			success: function($respon) {
				if ($respon.status == true) {
					$("#modal-tambah").modal('hide');
					successNotify($respon.message);
					$tabel1.ajax.reload(null, true);
				} else {
					errorNotify($respon.message);
				}
				enableButton("#form-tambah button[type=submit]")
			},
			error: function() {
				$("#form-tambah button[type=submit]").button('reset');
				$("#modal-tambah").modal('hide');
				errorNotify('Terjadi Kesalahan!');
			}
		});
		@endif


		@if(ucu())

		var quill_pendahuluan_edit = new Quill('#pendahuluan_edit', {
			placeholder: 'Tulis Salam Pembuka / Pendahuluan.....',
			theme: 'snow',
			modules: {
				toolbar: toolbarOptions
			}
		});
		quill_pendahuluan_edit.on('text-change', function(delta, oldDelta, source) {
			$("#form-edit #pendahuluan").val($("#pendahuluan_edit .ql-editor").html());
		});

		$validator_form_edit = $("#form-edit").validate();
		$("#modal-edit").on('show.bs.modal', function(e) {
			$uuid = $(e.relatedTarget).data('uuid');
			$validator_form_edit.resetForm();
			$("#form-edit").clearForm();
			disableButton("#form-edit button[type=submit]");
			$("#pendahuluan_edit .ql-editor").html('');
			$('#form-edit #jenis').selectize()[0].selectize.clear();
			$.get("{{url($main_path.'/get-data')}}/" + $uuid, function(respon) {
				if (respon.status) {
					$('#form-edit #kode').val(respon.data.kode);
					$('#form-edit #nama_sesi').val(respon.data.nama_sesi);
					$('#form-edit #jenis').selectize()[0].selectize.setValue(respon.data.jenis, false);		
					$('#form-edit #gambar').val(respon.data.gambar);
					$('#form-edit #uuid').val(respon.data.uuid);
					quill_pendahuluan_edit.clipboard.dangerouslyPasteHTML(respon.data.pendahuluan);
					$('#form-edit #pendahuluan').val(respon.data.pendahuluan);
					enableButton("#form-edit button[type=submit]");
				} else {
					errorNotify(respon.message);
				}
			})
		});

		$('#form-edit').ajaxForm({
			beforeSubmit: function() {
				disableButton("#form-edit button[type=submit]")
			},
			success: function($respon) {
				if ($respon.status == true) {
					$("#modal-edit").modal('hide');
					successNotify($respon.message);
					$tabel1.ajax.reload(null, true);
				} else {
					errorNotify($respon.message);
				}
				enableButton("#form-edit button[type=submit]")
			},
			error: function() {
				$("#form-edit button[type=submit]").button('reset');
				$("#modal-edit").modal('hide');
				errorNotify('Terjadi Kesalahan!');
			}
		});
		@endif
		

		@if(ucd())
		$('#form-delete').ajaxForm({
			beforeSubmit: function() {},
			success: function($respon) {
				if ($respon.status == true) {
					successNotify($respon.message);
					$tabel1.ajax.reload(null, true);
				} else {
					errorNotify($respon.message);
				}
			},
			error: function() {
				errorNotify('Terjadi Kesalahan!');
			}
		});
		var initKonfirmDelete = function() {
			$('.btn-konfirm-delete').on('click', function(e) {
				$uuid = $(this).data('uuid');

				$.get("{{url($main_path.'/get-data')}}/" + $uuid, function(respon) {
					if (respon.status) {
						$("#form-delete #uuid").val(respon.data.uuid);
						$.confirm({
							title: 'Yakin Hapus Data?',
							content: respon.informasi,
							buttons: {
								cancel: {
									text: 'Batalkan'
								},
								confirm: {
									text: 'Hapus',
									btnClass: 'btn-danger',
									action: function() {
										$("#form-delete").submit()
									}
								},
							}
						});
					} else {
						errorNotify(respon.message);
					}
				})
			})
		}
		@endif

		var initViewGambar = function(){

			$('.btn-view-image').on('click', function(e) {
				$filename = $(this).data('image');
				// $base_url_image = '{{url("gambar")}}';
				$base_url_image = '';
				if($filename){
		 			$.alert({
					    title: 'Gambar',
					    columnClass: 'col-md-6',
					    content: '<div width="100%"><center><img src="'+$filename+'" class="img-fluid rounded-lg" ></center></div>',
					});
			 	}
			})
		}

		var initViewUsing = function(){
			$('.btn-using').on('click', function(e) {
				$uuid = $(this).data('uuid');
				$("#panel-using").load("{{url($main_path)}}/using/" + $uuid, function(){
					$("#datatable-using").DataTable();
					$("#modal-using").modal('show');
				})
				
			});
		}
	})
</script>
@endsection