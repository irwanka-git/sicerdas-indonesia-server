<?php
$main_path = Request::segment(1);
loadHelper('akses,function'); 
$list_kelompok = get_list_enum_values('soal_mode_kerja_eng','kelompok');
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
				<h6 class="card-title" style="background-color: rgb(90, 56, 145); color:white; padding:5px; margin-bottom:10px;">ENGLISH VERSION</h6>
				<div class="card">
					<div class="card-header">
						<h5 class="card-title">{{$pagetitle}}</h5>
						<h6 class="card-subtitle text-muted">
                            {{$smalltitle}}
                         </h6>
					</div>
					<div class="card-body">
						@if(ucc())
				   		{{Html::btnModal('<i class="la la-plus-circle"></i> Tambah Soal','modal-tambah','primary')}}
				   		<hr>
				   		@endif
						<table id="datatable" class="table table-striped table-hover table-sm" style="width:100%">
							<thead>
								<tr>
									<th width="5%">#</th>									
									<th>Kelompok</th>
									<th>Pernyataan / Soal</th>
									<th>Pilihan Jawaban</th>
									<th>Nama Prioritas</th>
									<th>Deskripsi</th>
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
	{{Html::mOpenLG('modal-tambah','Tambah Soal')}}
    {{ Form::bsSelect2('Kelompok','kelompok',$list_kelompok,'',true,'md-8')}}
    {{ Form::bsNumeric('Nomor Urut','urutan','',true,'md-8') }}
	{{ Form::bsTextArea('Soal','soal','',true,'md-8') }}
	{{ Form::bsTextField('Aspek / Mode Kerja','mode_kerja','',true,'md-8') }}
	<p>Gunakan tanda <b>:</b> untuk memisahkan pilihan jawaban dan nama prioritas</p>
	{{ Form::bsTextField('Pilihan A : Prioritas A','pilihan_a','',true,'md-8') }}
	{{ Form::bsTextField('Pilihan B : Prioritas B','pilihan_b','',true,'md-8') }}
	{{ Form::bsTextField('Pilihan C : Prioritas C','pilihan_c','',true,'md-8') }}
	{{ Form::bsTextField('Pilihan D : Prioritas D','pilihan_d','',true,'md-8') }}
	{{ Form::bsTextField('Pilihan E : Prioritas E','pilihan_e','',true,'md-8') }}
	<div class="mb-3">
        <label class="form-label">Deskripsi <star>*</star> </label>
        <div  id="deskripsi_tambah"></div>
        <textarea style="display: none;" name="deskripsi" id="deskripsi"  required="required"></textarea>
    </div>

	{{Html::mCloseSubmitLG('Simpan')}}
{{ Form::bsClose()}}
@endif


@if(ucu())
 <!-- MODAL FORM TAMBAH -->
{{ Form::bsOpen('form-edit',url($main_path."/update")) }}
	{{Html::mOpenLG('modal-edit','Ubah Soal')}}
    {{ Form::bsSelect2('Kelompok','kelompok',$list_kelompok,'',true,'md-8')}}
	{{ Form::bsNumeric('Nomor Urut','urutan','',true,'md-8') }}
	{{ Form::bsTextArea('Soal','soal','',true,'md-8') }}
	{{ Form::bsTextField('Aspek / Mode Kerja','mode_kerja','',true,'md-8') }}
	<p>Gunakan tanda <b>:</b> untuk memisahkan pilihan jawaban dan nama prioritas</p>
	{{ Form::bsTextField('Pilihan A','pilihan_a','',true,'md-8') }}
	{{ Form::bsTextField('Pilihan B','pilihan_b','',true,'md-8') }}
	{{ Form::bsTextField('Pilihan C','pilihan_c','',true,'md-8') }}
	{{ Form::bsTextField('Pilihan D','pilihan_d','',true,'md-8') }}
	{{ Form::bsTextField('Pilihan E','pilihan_e','',true,'md-8') }}
    <div class="mb-3">
        <label class="form-label">deskripsi  <star>*</star> </label>
        <div  id="deskripsi_edit"></div>
        <textarea style="display: none;" name="deskripsi" id="deskripsi"  required="required"></textarea>
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


{{Html::mOpenLG('modal-lihat-soal','deskripsi')}}
	<div id="panel-lihat-soal">
		
	</div>
{{Html::mCloseLG()}}


@endsection

@section("js")
<script src="//cdn.quilljs.com/1.3.6/quill.min.js"></script>
<script type="text/javascript">
	$(function(){

			var toolbarOptions = [['bold', 'italic'], ['link', 'image']];
			var $tabel1 = $('#datatable').DataTable({
			    processing: true,
			    responsive: true,
			    fixedHeader: true,
			    serverSide: true,
			    ajax: "{{url($main_path.'/dt')}}",
			    "iDisplayLength": 25,
			    columns: [
			    	 {data:'DT_Row_Index' , orderable:false, searchable: false,sClass:""},
			         {data:'kelompok' , name:"kelompok" , orderable:false, searchable: false,sClass:""},
			         {data:'soal' , name:"soal" , orderable:false, searchable: false,sClass:""},
			         {data:'pilihan_jawaban' , name:"pilihan_jawaban" , orderable:false, searchable: false,sClass:""},
			         {data:'nama_prioritas' , name:"nama_prioritas" , orderable:false, searchable: false,sClass:""},
			         {data:'deskripsi' , name:"deskripsi" , orderable:false, searchable: false,sClass:""},
			         {data:'action' , orderable:false, searchable: false,sClass:"text-center"},
			        ],
			        "fnRowCallback": function( nRow, aData, iDisplayIndex, iDisplayIndexFull ) {
			        $(nRow).addClass( aData["rowClass"] );
			        return nRow;
			    },
			    "drawCallback": function( settings ) {
			        initKonfirmDelete();
			    }
			});

			$tabel1.on( 'responsive-display', function ( e, datatable, columns ) {
			    initKonfirmDelete();;
			});

			@if(ucc())

			 var quill_deskripsi_tambah = new Quill('#deskripsi_tambah', {
			    placeholder: 'Tulis Isi deskripsi.....',
			    theme: 'snow',
			    modules: {
				    toolbar: toolbarOptions
				  }
			  });
 			 quill_deskripsi_tambah.on('text-change', function(delta, oldDelta, source) {
 			 		$("#form-tambah #deskripsi").val($("#deskripsi_tambah .ql-editor").html());
 			 });

			$validator_form_tambah = $("#form-tambah").validate();
			$("#modal-tambah").on('show.bs.modal', function(e){
				$validator_form_tambah.resetForm();
				$("#form-tambah").clearForm(); 
				$("#deskripsi_tambah .ql-editor").html('');
                $('#form-tambah #kelompok').selectize()[0].selectize.clear();
				$("#form-tambah #deskripsi").val('');
				enableButton("#form-tambah button[type=submit]")
			});

			$('#form-tambah').ajaxForm({
				beforeSubmit:function(){
					disableButton("#form-tambah button[type=submit]")
				},
				success:function($respon){
					if ($respon.status==true){
						 $("#modal-tambah").modal('hide'); 
						 successNotify($respon.message);
						 $tabel1.ajax.reload(null, true);
					}else{
						errorNotify($respon.message);
					}
					enableButton("#form-tambah button[type=submit]")
				},
				error:function(){
					$("#form-tambah button[type=submit]").button('reset');
					$("#modal-tambah").modal('hide'); 
					errorNotify('Terjadi Kesalahan!');
				}
			}); 
			@endif

 		
		@if(ucu())

		var quill_deskripsi_edit = new Quill('#deskripsi_edit', {
		    placeholder: 'Tulis Isi deskripsi.....',
		    theme: 'snow',
		    modules: {
			    toolbar: toolbarOptions
			  }
		  });
			 quill_deskripsi_edit.on('text-change', function(delta, oldDelta, source) {
			 		$("#form-edit #deskripsi").val($("#deskripsi_edit .ql-editor").html());
			 });

		$validator_form_edit = $("#form-edit").validate();
		$("#modal-edit").on('show.bs.modal', function(e){
			$uuid  = $(e.relatedTarget).data('uuid');
			$validator_form_edit.resetForm();
			$("#form-edit").clearForm();
			disableButton("#form-edit button[type=submit]");
			$("#deskripsi_edit .ql-editor").html('');
			$.get("{{url($main_path.'/get-data')}}/"+$uuid, function(respon){
				if(respon.status){ 
					$('#form-edit #deskripsi').val(respon.data.deskripsi);
					$('#form-edit #urutan').val(respon.data.urutan);
					quill_deskripsi_edit.clipboard.dangerouslyPasteHTML(respon.data.deskripsi);				
					$('#form-edit #soal').val(respon.data.soal);
					$('#form-edit #mode_kerja').val(respon.data.mode_kerja);
					$('#form-edit #pilihan_a').val(respon.data.pilihan_a);
					$('#form-edit #pilihan_b').val(respon.data.pilihan_b);
					$('#form-edit #pilihan_c').val(respon.data.pilihan_c);
					$('#form-edit #pilihan_d').val(respon.data.pilihan_d);
					$('#form-edit #pilihan_e').val(respon.data.pilihan_e); 
                    $('#form-edit #kelompok').selectize()[0].selectize.setValue(respon.data.kelompok,false);
					$('#form-edit #uuid').val(respon.data.uuid);
					enableButton("#form-edit button[type=submit]");
				}else{
					errorNotify(respon.message);
				}
			})
		});

		$('#form-edit').ajaxForm({
			beforeSubmit:function(){disableButton("#form-edit button[type=submit]")},
			success:function($respon){
				if ($respon.status==true){
					 $("#modal-edit").modal('hide'); 
					 successNotify($respon.message);
					 $tabel1.ajax.reload(null, true);
				}else{
					errorNotify($respon.message);
				}
				enableButton("#form-edit button[type=submit]")
			},
			error:function(){
				$("#form-edit button[type=submit]").button('reset');
				$("#modal-edit").modal('hide'); 
				errorNotify('Terjadi Kesalahan!');
			}
		}); 
		@endif



		@if(ucd())
		$('#form-delete').ajaxForm({
			beforeSubmit:function(){},
			success:function($respon){
				if ($respon.status==true){
					 successNotify($respon.message);
					 $tabel1.ajax.reload(null, true);
				}else{
					errorNotify($respon.message);
				}
			},
			error:function(){errorNotify('Terjadi Kesalahan!');}
		}); 
		var initKonfirmDelete= function(){
			$('.btn-konfirm-delete').on('click', function(e){
				$uuid  = $(this).data('uuid');
				 
				$.get("{{url($main_path.'/get-data')}}/"+$uuid, function(respon){
					if(respon.status){
						$("#form-delete #uuid").val(respon.data.uuid);
						$.confirm({
						    title: 'Yakin Hapus Data?',
						    content: respon.informasi,
						    buttons: {
						        cancel :{
						        	text: 'Batalkan'
						        },
						        confirm: {
						        	text: 'Hapus',
						        	btnClass: 'btn-danger',
						        	action:function(){$("#form-delete").submit()}
						        },
						    }
						});
					}else{
						errorNotify(respon.message);
					}
				})
			})
		}
		@endif
	})
</script>
@endsection