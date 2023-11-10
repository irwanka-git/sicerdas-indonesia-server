<?php loadHelper('function');?>
<div class="row" style="font-size: 0.9em !important;">
    <div class="col-lg-12">
        <table id="datatable-using" class="table table-striped table-hover table-sm" style="width:100%">
            <thead>
                <tr>							
                    <th width="2%">#</th>
                    <th width="15%">Nama Tes</th>
                    <th width="10%">Tanggal</th>
                    <th width="10%">Lokasi</th>
                    <th width="10%">Pengguna</th>
                    <th width="15%">Kota</th>
                </tr>
            </thead>
            <tbody>
                <?php $no = 1;?>
                @foreach($data as $r)
                <tr>
                    <td><a target="_blank" href="{{url('/manajemen-sesi/detil/'.Crypt::encrypt($r->id_quiz))}}">{{$r->token}}</a></td>
                    <td>{{$r->nama_sesi}}</td>
                    <td>{{tgl_indo_singkat($r->tanggal)}}</td>
                    <td>{{$r->nama_lokasi}}</td>
                    <td>{{$r->nama_pengguna}}</td>
                    <td>{{$r->kota}}</td>
                </tr>
                @endforeach
            </tbody>
        </table>
    </div>
</div> 