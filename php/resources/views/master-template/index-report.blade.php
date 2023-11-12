<?php
// tabel: quiz_sesi_template
// quiz_sesi_template
// id_quiz_template int
// nama_sesi        varchar
// skoring_tabel    varchar
// uuid             char

$main_path = Request::segment(1);
loadHelper('akses,function');
$list_model = DB::select("select nama as text, id as value from model_report order by id asc ");

$list_report_query = DB::select("select nama_report as text, id_report as value , tabel_referensi,  tabel_terkait from quiz_sesi_report where jenis = 1 order by tabel_terkait desc, id_report asc");
$tabel_sesi = array();
foreach($sesi as $r){
    array_push($tabel_sesi, $r->tabel);
}

$list_model_report = array(); 
foreach ($list_model as $lm) {
	array_push($list_model_report, $lm);
}

$list_report = array();
$list_other= array();
foreach($list_report_query as $r){
	$explode_terkait = explode(",", $r->tabel_terkait);
	if (count($explode_terkait) > 1) {
		$valid  = true;
		foreach($explode_terkait as $ex) {
			if (!in_array($ex, $tabel_sesi)){
				$valid  = false;
			}
		}
		if ($valid){
			array_push($list_report, $r);
		}
	}else{
		if (in_array($r->tabel_terkait, $tabel_sesi)){
			array_push($list_report, $r);
		}
	}
    
    if($r->tabel_terkait =="-"){
        array_push($list_other, $r);
    }
}
$list_report = json_decode(json_encode($list_report));
 

$list_lampiran_query = DB::select("select nama_report as text, id_report as value , tabel_referensi, tabel_terkait from quiz_sesi_report where jenis = 2 order by tabel_terkait desc");

$list_lampiran = array();
foreach($list_lampiran_query as $r){
	$explode_terkait = explode(",", $r->tabel_terkait);
	if (count($explode_terkait) > 1) {
		$valid  = true;
		foreach($explode_terkait as $ex) {
			if (!in_array($ex, $tabel_sesi)){
				$valid  = false;
			}
		}
		if ($valid){
			array_push($list_lampiran, $r);
		}
	}else{
		if (in_array($r->tabel_terkait, $tabel_sesi)){
			array_push($list_lampiran, $r);
		}
	}
}
$list_lampiran = json_decode(json_encode($list_lampiran));

$list_kertas = json_decode(json_encode(array(["value"=>"A4", "text"=>"A4"], ["value"=>"Folio", "text"=>"Folio"])));
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
				<b class="pull-right">Kode : {{$template->kode}}-{{$template->id_quiz_template}} </b>
				<h5 class="card-title"><b>{{$template->nama_sesi}}</b></h5>
				<h6 class="card-subtitle text-muted">Fitur Ini Digunakan Untuk Manajemen Report Template Tes </h6>
			</div>
			<div class="card-body">
				<p>
					<a href="{{url('template-tes')}}" class="btn btn-secondary"> <i class="la la-arrow-left"></i> Kembali</a>
					@if(ucu())
					<a href="#" class="btn btn-warning pull-right" id="btn-create-dummy"> <i class="la la-cog"></i> Buat Data Dummy</a>
					<a href="#" class="btn btn-success" data-uuid="{{$template->uuid}}" data-bs-toggle="modal" data-bs-target="#modal-saran"> <i class="la la-edit"></i> Ubah Saran</a>
					<a href="#" class="btn btn-success" data-uuid="{{$template->uuid}}" data-bs-toggle="modal" data-bs-target="#modal-kertas"> <i class="la la-edit"></i> Ubah Ukuran Kertas ({{$template->kertas}})</a>
					@endif
				</p>
				<hr>
				<div class="row">
                    <div class="col-md-7">
						{{ Form::bsSelect2('Pilih Model Report','id_model_select',$list_model_report,'',false,'md-8')}} 
						<hr>
                        @if(ucc())
                        {{Html::btnModal('<i class="la la-plus-circle"></i> Komponen','modal-tambah','primary')}}
                        {{Html::btnModal('<i class="la la-plus-circle"></i> Lainnya','modal-tambah-lain','secondary')}}
						 @endif
                        <div id="komponen-laporan"></div>
						<hr>
						@if(ucc())
                        {{Html::btnModal('<i class="la la-plus-circle"></i> Tambah Lampiran','modal-lampiran','warning')}}
						 @endif
						<div id="lampiran-laporan"></div>
						<hr>
						<div>
							<button class="btn btn-outline-primary" href="#"  id="link-cetak"><i class="fa fa-download"></i> Download Sample PDF</button>
						</div>
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

{{ Form::bsOpen('form-export',url($main_path."/export-report-sample")) }}
	{{ Form::bsHidden('id_model',"") }} 
	{{ Form::bsHidden('id_quiz',"") }} 
	{{ Form::bsHidden('id_user',"") }} 
{{ Form::bsClose()}}

@if(ucc())
<!-- MODAL FORM TAMBAH -->
{{ Form::bsOpen('form-tambah',url($main_path."/insert-report")) }}
	{{Html::mOpenLG('modal-tambah','Tambah Komponen')}}
		{{ Form::bsHidden('id_quiz_template',$template->id_quiz_template) }} 
		{{ Form::bsHidden('model',"-") }} 
		{{ Form::bsSelect2('Komponen','id_report',$list_report,'',true,'md-8')}} 
	{{Html::mCloseSubmitLG('Simpan')}}
{{ Form::bsClose()}}
<!-- MODAL FORM TAMBAH -->
{{ Form::bsOpen('form-tambah-lain',url($main_path."/insert-report")) }}
	{{Html::mOpenLG('modal-tambah-lain','Tambah Lainnya')}}
		{{ Form::bsHidden('id_quiz_template',$template->id_quiz_template) }} 
		{{ Form::bsHidden('model',"-") }} 
		{{ Form::bsSelect2('Komponen','id_report',$list_other,'',true,'md-8')}} 
	{{Html::mCloseSubmitLG('Simpan')}}
{{ Form::bsClose()}}

<!-- MODAL FORM TAMBAH -->
{{ Form::bsOpen('form-lampiran',url($main_path."/insert-lampiran")) }}
	{{Html::mOpenLG('modal-lampiran','Tambah Lampiran')}}
		{{ Form::bsHidden('id_quiz_template',$template->id_quiz_template) }} 
		{{ Form::bsSelect2('Lampiran','id_report',$list_lampiran,'',true,'md-8')}} 
	{{Html::mCloseSubmitLG('Simpan')}}
{{ Form::bsClose()}}
@endif


@if(ucu())
<!-- MODAL FORM EDIT -->
{{ Form::bsOpen('form-saran',url($main_path."/update-saran")) }}
	{{Html::mOpenLG('modal-saran','Edit Saran / Rekomendasi')}}
		{{ Form::bsTextField('Judul','judul_saran','',true,'md-8') }}		
		<div class="mb-3">
			<label class="form-label">Isi <star>*</star> </label>
			<div  id="isi_edit"></div>
			<textarea style="display: none;" name="isi_saran" id="isi_saran"  required="required"></textarea>
		</div>
		{{ Form::bsHidden('id_quiz_template',$template->id_quiz_template) }} 
	{{Html::mCloseSubmitLG('Simpan')}}
{{ Form::bsClose()}}


{{ Form::bsOpen('form-kertas',url($main_path."/update-kertas")) }}
	{{Html::mOpenLG('modal-kertas','Ubah Ukuran Kertas')}}
		{{ Form::bsHidden('id_quiz_template',$template->id_quiz_template) }} 
		{{ Form::bsSelect2('Ukuran Kertas','kertas',$list_kertas,'',true,'md-8')}} 
	{{Html::mCloseSubmitLG('Simpan')}}
{{ Form::bsClose()}}

@endif

@if(ucd())
<!-- FORM DELETE -->
{{ Form::bsOpen('form-delete',url($main_path."/delete-report")) }}
	{{ Form::bsHidden('uuid','') }}
{{ Form::bsClose()}}

{{ Form::bsOpen('form-delete-lampiran',url($main_path."/delete-lampiran")) }}
	{{ Form::bsHidden('uuid','') }}
{{ Form::bsClose()}}


{{ Form::bsOpen('form-update',url($main_path."/update-urutan-report")) }}
    {{ Form::bsHidden('id_quiz_template',$template->id_quiz_template) }} 
    {{ Form::bsHidden('urutan_list','') }}    
{{ Form::bsClose()}}

{{ Form::bsOpen('form-update-lampiran',url($main_path."/update-urutan-lampiran")) }}
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
		$("#id_model_select").on('change', function(){
			reloadKomponen();
		});

        reloadKomponen = function(){
			$model = $("#id_model_select").val();
            $("#komponen-laporan").load("{{url($main_path)}}/komponen-report/{{$template->id_quiz_template}}/"+$model, function(){
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
				$("#form-tambah #model").val($model);
				$("#form-tambah-lain #model").val($model);
				$("#form-export #id_model").val($model);
            })
        }

		reloadLampiran= function(){
            $("#lampiran-laporan").load("{{url($main_path)}}/lampiran-report/{{$template->id_quiz_template}}", function(){
                $("#list-lampiran").sortable({
                    connectWith: "ul",
                    placeholder: "placeholder",
                    delay: 150,
                    stop: function(e, ui) {
                        $urutan = "";
                        $.map($(this).find('li'), function(el) {$urutan = $urutan + el.id + "," ;});
                        $("#form-update-lampiran #urutan_list").val($urutan);
                        $("#form-update-lampiran").submit();
                    }
                }).disableSelection();
                initKonfirmDeleteLampiran();
                initPreviewLampiran();
            })
        }

		cekDataDummyTemplate = function(){
			$.post("{{url($main_path)}}/cek-dummy-template/{{$template->id_quiz_template}}", {_token:"{{csrf_token()}}"}, function(respon){
				if (respon.status==false){
					confirmCreateDummy()
				}else{
					$("#form-export #id_quiz").val(respon.data.id_quiz);
					$("#form-export #id_user").val(respon.data.id_user);
				}
			})
		}

		confirmCreateDummy = function(){
			$.confirm({
					title: 'Buat data dummy',
					content: "Buat data pengujian (dummy) untuk jenis tes ini. ",
					buttons: {
						cancel: {
							text: 'Batalkan'
						},
						confirm: {
							text: 'Lanjutkan',
							btnClass: 'btn-success',
							action: function() {
								$.post("{{url($main_path)}}/create-dummy-template/{{$template->id_quiz_template}}", {_token:"{{csrf_token()}}"},  function(respon){
									if (respon.status==true){
										successNotify(respon.message);
										cekDataDummyTemplate();
									}
								})
							}
						},
					}
						});
		}

		$("#link-cetak").on("click", function(){
			$("#form-export").submit();
		})

		$('#form-export').ajaxForm({
			beforeSubmit: function() {
			},
			success: function(respon) {
				//  console.log(respon)
				if (respon.status == true){
					successNotify("Berhasil Buat Laporan Baru");
					window.open(respon.data, "_blank");
				}
			},
			error: function() {
			}
		});


		$("#btn-create-dummy").on('click', function(){
			confirmCreateDummy()
		})

        reloadKomponen();
        reloadLampiran();
		cekDataDummyTemplate();

		@if(ucc())

		var toolbarOptions = [['bold', 'italic'], ['link']];

		var quill_isi_saran = new Quill('#isi_edit', {
			    placeholder: 'Tulis Isi Saran.....',
			    theme: 'snow',
			    modules: {
				    toolbar: toolbarOptions
				  }
			  });

		 quill_isi_saran.on('text-change', function(delta, oldDelta, source) {
		 		$("#form-saran #isi_saran").val($("#isi_edit .ql-editor").html());
		 });
		 $validator_form_saran = $("#form-saran").validate();
		 $("#modal-saran").on('show.bs.modal', function(e) {
			$validator_form_saran.resetForm();
			$("#form-saran").clearForm(); 
			$('#form-saran #judul_saran').val('')
			$("#isi_edit .ql-editor").html('');
			disableButton("#form-saran button[type=submit]");
			$uuid = $(e.relatedTarget).data('uuid');
			$.get("{{url($main_path.'/get-data')}}/" + $uuid, function(respon) {
				if (respon.status) {
					$('#form-saran #judul_saran').val(respon.data.judul_saran);
					quill_isi_saran.clipboard.dangerouslyPasteHTML(respon.data.isi_saran);
					$('#form-saran #isi_saran').val(respon.data.isi_saran);
					enableButton("#form-saran button[type=submit]");
				} else {
					errorNotify(respon.message);
				}
			})
		});


		$('#form-saran').ajaxForm({
			beforeSubmit: function() {
				disableButton("#form-saran button[type=submit]")
			},
			success: function($respon) {
				if ($respon.status == true) {
					$("#modal-saran").modal('hide');
					successNotify($respon.message);
				} else {
					errorNotify($respon.message);
				}
				enableButton("#form-saran button[type=submit]")
			},
			error: function() {
				$("#form-saran button[type=submit]").button('reset');
				$("#modal-saran").modal('hide');
				errorNotify('Terjadi Kesalahan!');
			}
		});


		$("#modal-kertas").on('show.bs.modal', function(e) {
			$validator_form_saran.resetForm();
			$("#form-kertas").clearForm();  
			disableButton("#form-kertas button[type=submit]");
			$uuid = $(e.relatedTarget).data('uuid');
			$.get("{{url($main_path.'/get-data')}}/" + $uuid, function(respon) {
				if (respon.status) {
					$('#form-kertas #kertas').val(respon.data.kertas);
					enableButton("#form-kertas button[type=submit]");
				} else {
					errorNotify(respon.message);
				}
			})
		});


		$('#form-kertas').ajaxForm({
			beforeSubmit: function() {
				disableButton("#form-kertas button[type=submit]")
			},
			success: function($respon) {
				if ($respon.status == true) {
					$("#modal-kertas").modal('hide');
					successNotify($respon.message);
					window.location.reload();
				} else {
					errorNotify($respon.message);
				}
				enableButton("#form-kertas button[type=submit]")
			},
			error: function() {
				$("#form-kertas button[type=submit]").button('reset');
				$("#modal-kertas").modal('hide');
				errorNotify('Terjadi Kesalahan!');
			}
		});


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

		$validator_form_tambah_lain = $("#form-tambah-lain").validate();
		$("#modal-tambah-lain").on('show.bs.modal', function(e) {
			$validator_form_tambah_lain.resetForm();
			$("#form-tambah-lain").clearForm(); 
			$('#form-tambah-lain #id_report').selectize()[0].selectize.clear(); 
			enableButton("#form-tambah-lain button[type=submit]")
		});

		$('#form-tambah-lain').ajaxForm({
			beforeSubmit: function() {
				disableButton("#form-tambah-lain button[type=submit]")
			},
			success: function($respon) {
				if ($respon.status == true) {
					$("#modal-tambah-lain").modal('hide');
					successNotify($respon.message);
					reloadKomponen();
				} else {
					errorNotify($respon.message);
				}
				enableButton("#form-tambah-lain button[type=submit]")
			},
			error: function() {
				$("#form-tambah-lain button[type=submit]").button('reset');
				$("#modal-tambah-lain").modal('hide');
				errorNotify('Terjadi Kesalahan!');
			}
		});


		$validator_form_lampiran = $("#form-lampiran").validate();
		$("#modal-lampiran").on('show.bs.modal', function(e) {
			$validator_form_lampiran.resetForm();
			$("#form-lampiran").clearForm(); 
			$('#form-lampiran #id_report').selectize()[0].selectize.clear(); 
			enableButton("#form-lampiran button[type=submit]")
		});

		$('#form-lampiran').ajaxForm({
			beforeSubmit: function() {
				disableButton("#form-lampiran button[type=submit]")
			},
			success: function($respon) {
				if ($respon.status == true) {
					$("#modal-lampiran").modal('hide');
					successNotify($respon.message);
					reloadLampiran();
				} else {
					errorNotify($respon.message);
				}
				enableButton("#form-lampiran button[type=submit]")
			},
			error: function() {
				$("#form-lampiran button[type=submit]").button('reset');
				$("#modal-lampiran").modal('hide');
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


		var initKonfirmDeleteLampiran = function() {
			$('.btn-konfirm-delete-lampiran').on('click', function(e) {
				$uuid = $(this).data('uuid');
				$.get("{{url($main_path.'/get-data-lampiran')}}/" + $uuid, function(respon) {
					if (respon.status) {
						$("#form-delete-lampiran #uuid").val(respon.data.uuid);
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
										$("#form-delete-lampiran").submit()
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

		$('#form-update-lampiran').ajaxForm({
			beforeSubmit: function() {},
			success: function($respon) {
				if ($respon.status == true) {
					successNotify($respon.message);
					 reloadLampiran();
				} else {
					errorNotify($respon.message);
				}
			},
			error: function() {
				errorNotify('Terjadi Kesalahan!');
			}
		});

		$('#form-delete-lampiran').ajaxForm({
			beforeSubmit: function() {},
			success: function($respon) {
				if ($respon.status == true) {
					successNotify($respon.message);
					reloadLampiran();
				} else {
					errorNotify($respon.message);
				}
			},
			error: function() {
				errorNotify('Terjadi Kesalahan!');
			}
		});

		var initPreviewLampiran = function(){
			$('.btn-preview-lampiran').on('click', function(e) {
				$uuid = $(this).data('uuid')
				$("#preview").attr('src',"{{env('GO_URL').'/preview-lampiran-dummy'}}/" + $uuid)
				$('#modal-preview').modal('show'); 
			})
		}

		@endif
	})
</script>
@endsection