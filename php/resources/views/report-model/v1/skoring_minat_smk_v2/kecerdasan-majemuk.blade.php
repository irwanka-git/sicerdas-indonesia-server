<?php 
loadHelper('skoring,function');

?>
<?php
    $jumlah_pilihan = 5;
    $list_kecerdasan = DB::select("SELECT *  FROM ref_kecerdasan_majemuk");
    $kecerdasan = [];
    foreach($list_kecerdasan as $s){
        $kecerdasan[$s->no]['nama_kecil'] = $s->nama_kecil;
        $kecerdasan[$s->no]['nama'] = $s->nama_kecerdasan;
        $kecerdasan[$s->no]['icon'] = $s->icon;
    }
?>

<div  style="padding-left:5px;padding-right:5px;">
    <div class="box-title-003">
        <b>KECERDASAN MAJEMUK</b>
    </div> 
    <br>
    <div class="box-subtitle-003">
        <b>Disusun Berdasarkan Urutan Skala Kecerdasan Majemuk</b>
    </div>
</div>
<table width="100%" cellpadding="0" cellspacing="5">
    <tbody>
        <?php 
        for ($i=1;$i<=$jumlah_pilihan;$i++){
            $field_name = 'km_'.$i;
            $nomor = $data_skoring->$field_name;
        ?>
        <tr>
            <td width="5%" class="box-urutan-003">
                {{$i}}
            </td>
            <td width="10%" class="box-cell-icon2 box-icon-003" align="center" valign="align-middle">
                <img src="{{asset('images/icon/'.$kecerdasan[$nomor]['icon'])}}" height="90px">
            </td>
            <td width="85%" class="box-cell-info">
                <b>{{$kecerdasan[$nomor]['nama_kecil']}}</b> <small> - {{$kecerdasan[$nomor]['nama']}}</small>
            </td>
        </tr>
        <?php
        }
        ?>
    </tbody>
</table>
 