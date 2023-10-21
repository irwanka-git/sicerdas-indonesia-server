<?php 
loadHelper('skoring,function');
//kedinasan.png
?>
<?php
    $jumlah_pilihan = 3;
    $list_sekolah = DB::select("SELECT *  FROM ref_sekolah_dinas");
    $sekolah = [];
    foreach($list_sekolah as $s){
        $sekolah[$s->no]['akronim'] = $s->akronim;
        $sekolah[$s->no]['nama'] = $s->nama_sekolah_dinas;
    }
?>
<div  style="padding-left:5px;padding-right:5px;">
    <div class="box-title-005" style="border-left: solid 25px #c45424  !important;background: #D7835E !important;">
        <b>PEMINATAN SEKOLAH KEDINASAN</b>
    </div> 
    <br>
    <div class="box-subtitle-005" style="border-left: solid 25px #c45424  !important;background: #D7835E !important;">
         Keterangan: Urutan nomor 1 (satu)  sampai 3 (tiga) yang paling dominan dari 12 (dua belas) pilihan sekolah kedinasan 
    </div>
</div>

<table width="100%" cellpadding="0" cellspacing="5">
    <tbody>
        <?php 
        for ($i=1;$i<=$jumlah_pilihan;$i++){
            $field_name = 'minat_dinas'.$i;
            $nomor = $data_skoring->$field_name;
        ?>
        <tr>
            <td width="5%" class="box-urutan-205">
                {{$i}}
            </td>
            <td width="10%" class="box-cell-icon2 box-icon-205" align="center" valign="align-middle">
                <img src="{{asset('images/icon/kedinasan.png')}}" height="60px">
            </td>
            <td width="85%" class="box-cell-info">
                <b>{{$sekolah[$nomor]['akronim']}}</b> <small> - {{$sekolah[$nomor]['nama']}}</small>
            </td>
        <tr>
        <?php
        }
        ?>
    </tbody>
</table>
