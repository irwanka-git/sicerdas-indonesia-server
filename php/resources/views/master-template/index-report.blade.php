<?php
// tabel: quiz_sesi_template
// quiz_sesi_template
// id_quiz_template int
// nama_sesi        varchar
// skoring_tabel    varchar
// uuid             char

$main_path = Request::segment(1);
loadHelper('akses,function');
$list_report_query = DB::select("select nama_report as text, id_report as value , tabel_referensi from quiz_sesi_report");
$tabel_sesi = array();
foreach($sesi as $r){
    array_push($tabel_sesi, $r->tabel);
}

$list_report = array();
foreach($list_report_query as $r){
    if (in_array($r->tabel_referensi, $tabel_sesi)){
        array_push($list_report, $r);
    }

    if($r->tabel_referensi =="-"){
        array_push($list_report, $r);
    }
}
$list_report = json_decode(json_encode($list_report));
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
				<h5 class="card-title"><b>{{$template->nama_sesi}}</b></h5>
				<h6 class="card-subtitle text-muted">Fitur Ini Digunakan Untuk Manajemen Report Template Tes </h6>
			</div>
			<div class="card-body">
				<a href="{{url('template-tes')}}" class="btn btn-secondary"> <i class="la la-arrow-left"></i> Kembali</a>
				<hr>
				<div class="row">
                    
                    <div class="col-md-7">
                        @if(ucc())
                        {{Html::btnModal('<i class="la la-plus-circle"></i> Tambah Komponen','modal-tambah','primary')}}
                        @endif
                        <div id="komponen-laporan"></div>
						<div id="lampiran-laporan"></div>
                    </div>
                    <div class="col-md-5">
                        <b>Daftar Sesi Tes: </b>
                         <ul>
                            @foreach ($sesi as $item)
                            <li>{{$item->urutan}}. {{$item->nama_sesi_ujian}}</li>
                            @endforeach
                         </ul>
                    </div>
                </div>
			</div>
		</div>
	</div>
</div>

@endsection

@section("modal")

@if(ucc())
<!-- MODAL FORM TAMBAH -->
{{ Form::bsOpen('form-tambah',url($main_path."/insert-report")) }}
	{{Html::mOpenLG('modal-tambah','Tambah Komponen')}}
		{{ Form::bsHidden('id_quiz_template',$template->id_quiz_template) }} 
		{{ Form::bsSelect2('Komponen','id_report',$list_report,'',true,'md-8')}} 
	{{Html::mCloseSubmitLG('Simpan')}}
{{ Form::bsClose()}}
@endif

 

@if(ucd())
<!-- FORM DELETE -->
{{ Form::bsOpen('form-delete',url($main_path."/delete-report")) }}
	{{ Form::bsHidden('uuid','') }}
{{ Form::bsClose()}}

{{ Form::bsOpen('form-update',url($main_path."/update-urutan-report")) }}
    {{ Form::bsHidden('id_quiz_template',$template->id_quiz_template) }} 
    {{ Form::bsHidden('urutan_list','') }}    
{{ Form::bsClose()}}
@endif

{{Html::mOpenLG('modal-preview','Preview Komponen')}}
	<div >
		<iframe id="preview" src="url"  height="400px" width="100%" title="Preview Komponen"></iframe>
	</div>
{{Html::mCloseLG()}}
@endsection

@section("js")
<script src="//cdn.quilljs.com/1.3.6/quill.min.js"></script>
<script src="https://ajax.googleapis.com/ajax/libs/jqueryui/1.11.2/jquery-ui.min.js"></script>
<script type="text/javascript">
	$(function() {
        reloadKomponen = function(){
            $("#komponen-laporan").load("{{url($main_path)}}/komponen-report/{{$template->id_quiz_template}}", function(){
                $("#list-report").sortable({
                    connectWith: "ul",
                    placeholder: "placeholder",
                    delay: 150,
                    stop: function(e, ui) {
                        $urutan = "";
                        $.map($(this).find('li'), function(el) {$urutan = $urutan + el.id + "," ;});
                        $("#form-update #urutan_list").val($urutan);
                        $("#form-update").submit();
                    }
                }).disableSelection();
                initKonfirmDelete();
                initPreview();
            })
        }
        reloadKomponen();
		@if(ucc())
		$validator_form_tambah = $("#form-tambah").validate();
		$("#modal-tambah").on('show.bs.modal', function(e) {
			$validator_form_tambah.resetForm();
			$("#form-tambah").clearForm(); 
			$('#form-tambah #id_report').selectize()[0].selectize.clear(); 
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
					reloadKomponen();
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

 

		@if(ucd())

        $('#form-update').ajaxForm({
			beforeSubmit: function() {},
			success: function($respon) {
				if ($respon.status == true) {
					successNotify($respon.message);
					 reloadKomponen();
				} else {
					errorNotify($respon.message);
				}
			},
			error: function() {
				errorNotify('Terjadi Kesalahan!');
			}
		});

		$('#form-delete').ajaxForm({
			beforeSubmit: function() {},
			success: function($respon) {
				if ($respon.status == true) {
					successNotify($respon.message);
					reloadKomponen();
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
				$.get("{{url($main_path.'/get-data-report')}}/" + $uuid, function(respon) {
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

		var initPreview = function(){
			$('.btn-preview-komponen').on('click', function(e) {
				$uuid = $(this).data('uuid')
				$("#preview").attr('src',"{{env('GO_URL').'/preview-report-dummy'}}/" + $uuid)
				$('#modal-preview').modal('show'); 
			})
		}
		@endif
	})
</script>
@endsection