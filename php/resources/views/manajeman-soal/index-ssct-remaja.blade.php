<?php
$main_path = Request::segment(1);
loadHelper('akses,function'); 
$list_komponen = get_list_enum_values('soal_ssct_remaja','komponen');
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
						<h5 class="card-title">Soal Subjek Penialain SSCT Remaja</h5>
						<h6 class="card-subtitle text-muted">Fitur Ini Digunakan Untuk Input Soal</h6>
					</div>
					<div class="card-body">
						@if(ucc())
				   		{{Html::btnModal('<i class="la la-plus-circle"></i> Tambah Subjek','modal-tambah','primary')}}
				   		<hr>
				   		@endif
						<table id="datatable" class="table table-striped table-hover table-sm" style="width:100%">
							<thead>
								<tr>
									<th width="5%">No.</th>
									<th width="15%">Komponen</th>
									<th width="20%">Aspek</th> 
									<th>Subjek</th>
									<th>Sikap Negatif</th>
									<th>Sikap Positif</th>
									<th width="5%">Actions</th>
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
	{{Html::mOpenLG('modal-tambah','Tambah Pelajaran')}}

		{{ Form::bsNumeric('Urutan','urutan','',true,'md-8') }}
		{{ Form::bsSelect2('Komponen','komponen',$list_komponen,'',true,'md-8')}}
		{{ Form::bsTextField('Subjek Penilaian','subjek_penilaian','',true,'md-8') }}
		{{ Form::bsTextField('Aspek','aspek','',true,'md-8') }} 
		<hr>
		<div class="row">
			<div class="col-md-6">
				{{ Form::bsTextField('Sikap Negatif 1','sikap_negatif1','',true,'md-8') }}
			</div>
			<div class="col-md-6">
			{{ Form::bsTextField('Sikap Positif 1','sikap_positif1','',true,'md-8') }}
			</div>
		</div>
		<hr>
		<div class="row">
			<div class="col-md-6">
				{{ Form::bsTextField('Sikap Negatif 2','sikap_negatif2','',true,'md-8') }}
			</div>
			<div class="col-md-6">
			{{ Form::bsTextField('Sikap Positif 2','sikap_positif2','',true,'md-8') }}
			</div>
		</div>
		<hr>
		<div class="row">
			<div class="col-md-6">
				{{ Form::bsTextField('Sikap Negatif 3','sikap_negatif3','',true,'md-8') }}
			</div>
			<div class="col-md-6">
			{{ Form::bsTextField('Sikap Positif 3','sikap_positif3','',true,'md-8') }}
			</div>
		</div>
	{{Html::mCloseSubmitLG('Simpan')}}
{{ Form::bsClose()}}
@endif


@if(ucu())
 <!-- MODAL FORM TAMBAH -->
{{ Form::bsOpen('form-edit',url($main_path."/update")) }}
	{{Html::mOpenLG('modal-edit','Ubah Soal')}}

        {{ Form::bsNumeric('Urutan','urutan','',true,'md-8') }}
        {{ Form::bsSelect2('Komponen','komponen',$list_komponen,'',true,'md-8')}}
        {{ Form::bsTextField('Subjek Penilaian','subjek_penilaian','',true,'md-8') }}
        {{ Form::bsTextField('Aspek','aspek','',true,'md-8') }} 
		<hr>
		<div class="row">
			<div class="col-md-6">
				{{ Form::bsTextField('Sikap Negatif 1','sikap_negatif1','',true,'md-8') }}
			</div>
			<div class="col-md-6">
			{{ Form::bsTextField('Sikap Positif 1','sikap_positif1','',true,'md-8') }}
			</div>
		</div>
		<hr>
		<div class="row">
			<div class="col-md-6">
				{{ Form::bsTextField('Sikap Negatif 2','sikap_negatif2','',true,'md-8') }}
			</div>
			<div class="col-md-6">
			{{ Form::bsTextField('Sikap Positif 2','sikap_positif2','',true,'md-8') }}
			</div>
		</div>
		<hr>
		<div class="row">
			<div class="col-md-6">
				{{ Form::bsTextField('Sikap Negatif 3','sikap_negatif3','',true,'md-8') }}
			</div>
			<div class="col-md-6">
			{{ Form::bsTextField('Sikap Positif 3','sikap_positif3','',true,'md-8') }}
			</div>
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


{{Html::mOpenLG('modal-lihat-soal','Pelajaran')}}
	<div id="panel-lihat-soal">
		
	</div>
{{Html::mCloseLG()}}

 

@endsection

@section("js")
<script src="//cdn.quilljs.com/1.3.6/quill.min.js"></script>
<script type="text/javascript">
	$(function(){

			var toolbarOptions = [['bold', 'italic'], ['link', 'image'],[{'list': 'bullet'}]];
			var $tabel1 = $('#datatable').DataTable({
			    processing: true,
			    responsive: true,
			    fixedHeader: true,
			    serverSide: true,
			    ajax: "{{url('soal-ssct-remaja/dt')}}",
			    "iDisplayLength": 25,
			    columns: [
			         {data:'urutan' , name:"urutan" , orderable:false, searchable: false,sClass:"text-center"},
			         {data:'komponen' , name:"komponen" , orderable:true, searchable: false,sClass:""},
			         {data:'aspek' , name:"aspek" , orderable:true, searchable: false,sClass:""}, 
			         {data:'subjek_penilaian' , name:"subjek_penilaian" , orderable:false, searchable: false,sClass:""},
                     {data:'sikap_negatif' , name:"sikap_negatif" , orderable:false, searchable: false,sClass:""},
                     {data:'sikap_positif' , name:"sikap_positif" , orderable:false, searchable: false,sClass:""},
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

			 
			$validator_form_tambah = $("#form-tambah").validate();
			$("#modal-tambah").on('show.bs.modal', function(e){
				$validator_form_tambah.resetForm();
				$("#form-tambah").clearForm();
				$("#form-tambah #subjek_penilaian").val('');
				$('#form-tambah #komponen').selectize()[0].selectize.clear();
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

		 
		$validator_form_edit = $("#form-edit").validate();
		$("#modal-edit").on('show.bs.modal', function(e){
			$uuid  = $(e.relatedTarget).data('uuid');
			$validator_form_edit.resetForm();
			$('#form-edit #komponen').selectize()[0].selectize.clear();
			$("#form-edit").clearForm();
			disableButton("#form-edit button[type=submit]");
			$.get("{{url('soal-ssct-remaja/get-data')}}/"+$uuid, function(respon){
				if(respon.status){
					$('#form-edit #subjek_penilaian').val(respon.data.subjek_penilaian);
					$('#form-edit #aspek').val(respon.data.aspek); 
					$('#form-edit #komponen').selectize()[0].selectize.setValue(respon.data.komponen,false);
					$('#form-edit #sikap_negatif1').val(respon.data.sikap_negatif1);
					$('#form-edit #sikap_positif1').val(respon.data.sikap_positif1);
					$('#form-edit #sikap_negatif2').val(respon.data.sikap_negatif2);
					$('#form-edit #sikap_positif2').val(respon.data.sikap_positif2);
					$('#form-edit #sikap_negatif3').val(respon.data.sikap_negatif3);
					$('#form-edit #sikap_positif3').val(respon.data.sikap_positif3);
					$('#form-edit #urutan').val(respon.data.urutan);
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
				 
				$.get("{{url('soal-sikap-pelajaran-kuliah/get-data')}}/"+$uuid, function(respon){
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