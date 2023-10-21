@foreach($list_lokasi as $r)
						<div class="col-12 col-md-6 col-lg-3">


							<div class="card">
								<div class="card-header px-4 pt-4">
									<?php 
									//$url_detil = url('/manajemen-sesi/explore/'.$r->tahun);
									$url_detil = url('/manajemen-sesi/explore/'.$data_tahun['tahun'].'/'.$data_biro['id_user_biro'].'/'.$data_provinsi->id.'/'.$r->id_lokasi);
									?>
									<div class="badge bg-primary my-2"><i class="las la-map-marker-alt"></i> {{$r->kabupaten}}</div>
									<a href="{{$url_detil}}"><h5 class="card-title mb-0"><i class="las la-university"></i> {{$r->nama_lokasi}}</h5></a>
									<p class="mb-2 fw-bold">Jumlah Tes<span class="float-end">{{$r->jumlah_tes}}</span></p>
								</div>

							</div>
						</div>
						@endforeach