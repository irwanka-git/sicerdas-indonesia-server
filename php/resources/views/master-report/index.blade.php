<?php
// tabel: quiz_sesi_master
// id_sesi_master		int
// kategori				char
// nama_sesi_ujian		char
// metode_skoring		char
// mode					char
// jawaban				int
// petunjuk_sesi		varchar

$main_path = Request::segment(1);
loadHelper('akses,function');
$list_tabel  =  DB::select("select tabel as value, tabel as text from quiz_sesi_master where tabel !='' group by tabel");
// dd($list_tabel);
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
				<h5 class="card-title">Master Report</h5>
				<h6 class="card-subtitle text-muted">Fitur Ini Digunakan Untuk Manajemen Data Master Report Tes Si-Cerdas </h6>
			</div>
			<div class="card-body">
				@if(ucc())
				{{Html::btnModal('<i class="la la-plus-circle"></i> Tambah Report','modal-tambah','primary')}}
				<hr>
				@endif
				<table id="datatable" class="table table-striped table-hover table-sm" style="width:100%">
					<thead>
						<tr>							
							<th width="5%">ID sesi</th> 
							<th width="20%">Nama Report</th>
							<th width="15%">Tabel Skoring</th>
                            <th width="15%">Blade Report</th>
							<th width="10%">Actions</th>
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
	{{Html::mOpenLG('modal-tambah','Tambah Report')}}
		{{ Form::bsTextField('Nama Report','nama','',true,'md-8') }}
		{{ Form::bsSelect2('Tabel Skoring','tabel_skoring',$list_tabel,'',true,'md-8')}}  
		{{ Form::bsTextField('Blade','blade','',true,'md-8') }}
	{{Html::mCloseSubmitLG('Simpan')}}
{{ Form::bsClose()}}
@endif


@if(ucu())
<!-- MODAL FORM EDIT -->
{{ Form::bsOpen('form-edit',url($main_path."/update")) }}
	{{Html::mOpenLG('modal-edit','Edit Master Sesi')}}
		{{ Form::bsTextField('Kategori','kategori','',true,'md-8') }}
		{{ Form::bsTextField('Nama Sesi','nama_sesi_ujian','',true,'md-8') }}
		{{ Form::bsTextField('Soal URL (Route)','soal','',true,'md-8') }}
		{{ Form::bsTextField('Tabel (skoring)','tabel','',true,'md-8') }}
		{{ Form::bsSelect2('Mode','mode',$list_tabel,'',true,'md-8')}}
		{{ Form::bsNumeric('Jawaban','jawaban','',true,'md-4') }}
		{{ Form::bsNumeric('Jumlah Karakter Jawaban','panjang_jawaban','',true,'md-4') }}
		<div class="mb-3">
			<label class="form-label">Petunjuk Sesi  <star>*</star> </label>
			<div  id="petunjuk_edit"></div>
			<textarea style="display: none;" name="petunjuk_sesi" id="petunjuk_sesi"  required="required"></textarea>
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


@endsection

@section("js")
<script src="//cdn.quilljs.com/1.3.6/quill.min.js"></script>
<script type="text/javascript">
	$(function() {

		var toolbarOptions = [['bold', 'italic'],['link', 'image']];
		var $tabel1 = $('#datatable').DataTable({
			processing: true,
			responsive: true,
			fixedHeader: true,
			serverSide: true,
			ajax: "{{url('master-report/dt')}}",
			"iDisplayLength": 25,
			columns: [
				{data: 'id_sesi_master',name: "id_sesi_master",orderable: false,searchable: false,sClass: "text-center"},
				{data: 'kategori',name: "kategori",orderable: false,searchable: false,sClass: ""},
				{data: 'nama_sesi_ujian',name: "nama_sesi_ujian",orderable: false,searchable: false,sClass: ""},
				{data: 'tabel',name: "tabel",orderable: false,searchable: false,sClass: "text-center"},
				{data: 'mode',name: "mode",orderable: false,searchable: false,sClass: "text-center"},
				{data: 'jawaban',name: "jawaban",orderable: false,searchable: false,sClass: "text-center"},
				{data: 'panjang_jawaban',name: "panjang_jawaban",orderable: false,searchable: false,sClass: "text-center"},
				{data: 'action',orderable: false,searchable: false,sClass: "text-center"},
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
			}
		});

		$tabel1.on('responsive-display', function(e, datatable, columns) {
			initKonfirmDelete();;
		});

	
		@if(ucc())

		var quill_petunjuk_tambah = new Quill('#petunjuk_tambah', {
			placeholder: 'Tulis Isi Petunjuk.....',
			theme: 'snow',
			modules: {
				toolbar: toolbarOptions
			}
		});
		quill_petunjuk_tambah.on('text-change', function(delta, oldDelta, source) {
			$("#form-tambah #petunjuk_sesi").val($("#petunjuk_tambah .ql-editor").html());
		});

		$validator_form_tambah = $("#form-tambah").validate();
		$("#modal-tambah").on('show.bs.modal', function(e) {
			$validator_form_tambah.resetForm();
			$("#form-tambah").clearForm();
			$("#petunjuk_tambah .ql-editor").html('');
			$("#form-tambah #petunjuk_sesi").val('');
			$('#form-tambah #mode').selectize()[0].selectize.clear();
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

		var quill_petunjuk_edit = new Quill('#petunjuk_edit', {
			placeholder: 'Tulis Isi Petunjuk Sesi.....',
			theme: 'snow',
			modules: {
				toolbar: toolbarOptions
			}
		});
		quill_petunjuk_edit.on('text-change', function(delta, oldDelta, source) {
			$("#form-edit #petunjuk_sesi").val($("#petunjuk_edit .ql-editor").html());
		});

		$validator_form_edit = $("#form-edit").validate();
		$("#modal-edit").on('show.bs.modal', function(e) {
			$uuid = $(e.relatedTarget).data('uuid');
			$validator_form_edit.resetForm();
			$("#form-edit").clearForm();
			disableButton("#form-edit button[type=submit]");
			$("#petunjuk_edit .ql-editor").html('');
			$('#form-edit #mode').selectize()[0].selectize.clear(); 
			$.get("{{url('master-report/get-data')}}/" + $uuid, function(respon) {
				if (respon.status) {
					
					$('#form-edit #kategori').val(respon.data.kategori);
					$('#form-edit #nama_sesi_ujian').val(respon.data.nama_sesi_ujian);
					$('#form-edit #soal').val(respon.data.soal);
					$('#form-edit #tabel').val(respon.data.tabel); 
					$('#form-edit #mode').selectize()[0].selectize.setValue(respon.data.mode, false);				
					$('#form-edit #jawaban').val(respon.data.jawaban);
					$('#form-edit #panjang_jawaban').val(respon.data.panjang_jawaban);
					quill_petunjuk_edit.clipboard.dangerouslyPasteHTML(respon.data.petunjuk_sesi);
					$('#form-edit #petunjuk_sesi').val(respon.data.petunjuk_sesi);
					
					$('#form-edit #uuid').val(respon.data.uuid);
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

				$.get("{{url('master-report/get-data')}}/" + $uuid, function(respon) {
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
	})
</script>
@endsection