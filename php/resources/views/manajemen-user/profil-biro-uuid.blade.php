<?php
$main_path = Request::segment(1);
loadHelper('akses'); 
$user = DB::table('users')->where('uuid', $uuid)->first();
?>
@extends('layout')
@section("pagetitle")
	 
@endsection

@section('content')

		<div class="row">
			<div class="col-12">
				<div class="card">
					<div class="card-header">
						<h5 class="card-title">Pengaturan Profil Biro</h5>
						<h6 class="card-subtitle text-muted">Fitur Ini Digunakan Untuk Pengaturan Data Profil Biro </h6>
					</div>
					<div class="card-body">
						{{ Form::bsOpen('form-profil',url("manajemen-akun-biro/update-profil")) }}
						<div class="row">
							<div class="col-md-6 pr-2">
								{{ Form::bsTextField('Nama Biro','nama_pengguna',$user->nama_pengguna,true,'md-8') }}
		                        {{ Form::bsTextArea('Alamat','alamat',$user->alamat,true,'md-8') }}
		                        {{ Form::bsTextField('Email','email',$user->email,true,'md-8') }}
		                        {{ Form::bsTextField('Nomor Telepon','telp',$user->telp,true,'md-8') }}
							</div>
							<div class="col-md-6">
								<p>Kop Biro<br>
		                    	<small>
									Ukuran Gambar yang Disarankan adalah 800 x 90 Pixels |
								<a target="_blank" href="https://www.canva.com/design/DAFzp7SUnNs/NkDtRV9GAQw6YAqD0zX8Jw/edit?utm_content=DAFzp7SUnNs&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton">Gunakan Template Canva</a>
								</small><br>
									<button id="btn_gambar" class="btn btn-secondary btn-upload-gambar" type="button">
		                    		<i class="la la-upload"></i> Upload Kop</button> 
		                    	</p>
								
		                    	{{ Form::bsHidden('kop_biro',$user->kop_biro) }}
		                    	<p id="gambar-kop">
		                    		@if($user->kop_biro)
		                    		<img class="img img-responsive img-thumbnail" src="{{url('kop/'.$user->kop_biro)}}">
		                    		@endif
		                    	</p>
								<hr>
								<p>Cover Biro (Format PDF) Sicerdas Versi 1<br>
									<small>File Sesuai Template | <a target="_blank" href="https://docs.google.com/document/d/1_8IzxLI0UVkaURLlbCRthhOalBO9QjP5/edit?usp=sharing&ouid=113639986592277865781&rtpof=true&sd=true">Download Template</a></small><br>
									<button id="btn_cover" class="btn btn-secondary btn-upload-gambar" type="button">
										<i class="la la-upload"></i> Upload Cover</button> 
								</p>
								{{ Form::bsHidden('cover_biro',$user->cover_biro) }}
								<p id="lihat-cover">
		                    		@if($user->cover_biro)
		                    		<a target="_blank" href="{{url('cover/'.$user->cover_biro)}}">Lihat Cover PDF</a>
		                    		@endif
		                    	</p>
								<hr>
		                    	<p>Cover Biro (Format Gambar) Sicerdas Versi 2<br>
								<small>Ukuran Gambar yang Disarankan adalah 1273 × 2000 Pixels | <a target="_blank" href="https://www.canva.com/design/DAFzyNDLdg0/ql3EMWELvf_zCoSX4-tUtg/view">Gunakan Template Canva</a></small><br>
		                    	<button id="btn_cover_gambar" class="btn btn-secondary btn-upload-gambar" type="button">
		                    		<i class="la la-upload"></i> Upload Cover</button> 
		                    	</p>
								
		                    	{{ Form::bsHidden('cover_biro_gambar',$user->cover_biro_gambar) }}
		                    	{{ Form::bsHidden('uuid',$user->uuid) }}
		                    	<p id="lihat-cover-gambar">
		                    		@if($user->cover_biro_gambar)
		                    		<a target="_blank" href="{{url('cover/'.$user->cover_biro_gambar)}}">Lihat Cover Gambar</a>
		                    		@endif
		                    	</p>
		                    	<hr>
		                    	 <div class="float-end w100"> 
			                          	<button type="submit" id="btn-submit" class="btn btn-primary"><i class="la la-save"></i> Simpan Perubahan</button>
			                          </div>
							</div>
						</div>
                    	{{ Form::bsClose()}}
					</div>
				</div>
			</div>
		</div>
 
@endsection

@section("modal")
 

{{ Form::bsOpen('form-upload-gambar',url($main_path."/upload-gambar")) }}
	 <input type="file" style="display: none;" id="upload-gambar" name="image" accept="image/*">
{{ Form::bsClose()}}

{{ Form::bsOpen('form-upload-cover',url($main_path."/upload-cover")) }}
	 <input type="file" style="display: none;" id="upload-cover" name="cover" accept=".pdf">
{{ Form::bsClose()}}

{{ Form::bsOpen('form-upload-cover-gambar',url($main_path."/upload-cover-gambar")) }}
	 <input type="file" style="display: none;" id="upload-cover-gambar" name="cover" accept="image/*">
{{ Form::bsClose()}}

@endsection

@section("js")
<script type="text/javascript">
	$(function(){
		$("#upload-gambar").on('change', function(){
	 		if($(this).val()){
	 			$("#form-upload-gambar").submit();
	 		}
		});

	 	$('#form-upload-gambar').ajaxForm({
			beforeSubmit:function(){ disableButton("#profil #btn-submit") },
			success:function($respon){
				//enableButton("#"+$form_gambar +" #btn_"+$field_gambar);
				//console.log($respon);
				if($respon.status==true){
					$("#kop_biro").val($respon.filename)
					$("#upload-gambar").val('');
					$("#gambar-kop").html("<img class='img img-responsive img-thumbnail' src='"+$respon.url_image+"'>");
				}else{
					errorNotify($respon.message);
				}
				
			},
			error:function(){errorNotify('Terjadi Kesalahan!');$("#upload-gambar").val('');}
		}); 

		$("#btn_gambar").on('click', function(){
			$("#upload-gambar").trigger('click');
		})


		$("#upload-cover").on('change', function(){
	 		if($(this).val()){
	 			$("#form-upload-cover").submit();
	 		}
		});
		

	 	$('#form-upload-cover').ajaxForm({
			beforeSubmit:function(){ disableButton("#profil #btn-submit") },
			success:function($respon){
				//enableButton("#"+$form_gambar +" #btn_"+$field_gambar);
				//console.log($respon);
				if($respon.status==true){
					$("#cover_biro").val($respon.filename)
					$("#upload-cover").val('');
					$("#lihat-cover").html("<a target='_blank' href='"+$respon.url+"'>Lihat Cover</a>");
				}else{
					errorNotify($respon.message);
				}
				
			},
			error:function(){errorNotify('Terjadi Kesalahan!');$("#upload-gambar").val('');}
		}); 

		$("#upload-cover-gambar").on('change', function(){
	 		if($(this).val()){
	 			$("#form-upload-cover-gambar").submit();
	 		}
		});

		$('#form-upload-cover-gambar').ajaxForm({
			beforeSubmit:function(){ disableButton("#profil #btn-submit") },
			success:function($respon){
				//enableButton("#"+$form_gambar +" #btn_"+$field_gambar);
				//console.log($respon);
				if($respon.status==true){
					$("#cover_biro_gambar").val($respon.filename)
					$("#upload-cover-gambar").val('');
					$("#lihat-cover-gambar").html("<a target='_blank' href='"+$respon.url+"'>Lihat Cover</a>");
				}else{
					errorNotify($respon.message);
				}
			},
			error:function(){errorNotify('Terjadi Kesalahan!');$("#upload-gambar").val('');}
		}); 

		$("#btn_cover").on('click', function(){
			$("#upload-cover").trigger('click');
		})

		$("#btn_cover_gambar").on('click', function(){
			$("#upload-cover-gambar").trigger('click');
		})


		$('#form-profil').ajaxForm({
			beforeSubmit:function(){disableButton("#form-profil button[type=submit]")},
			success:function($respon){
				if ($respon.status==true){
					 successNotify($respon.message);
				}else{
					errorNotify($respon.message);
				}
				enableButton("#form-profil button[type=submit]")
			},
			error:function(){
				$("#form-profil button[type=submit]").button('reset');
				errorNotify('Terjadi Kesalahan!');
			}
		}); 
	})
</script>
@endsection