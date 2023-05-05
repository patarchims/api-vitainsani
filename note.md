// ANTREAN RESEP


req_type => online, ofline,
jenis_pasien => debitur,




<!-- jenis pasien -->

MRC 
NMRC


kode_booking_ref => di dapat dari antrian_ol kode_booking
mrn => nomor rekam medik
cdttm => time booking,
jenis => online
jenis => Racikan & Non racikan => nrc
tanggal => tanggal insert hari H
jam => jam saat ini
time_elapsed => berapa lama dilayani , yang statusnya false,
1 nomor 10 menit, => /tanggal hari ini.
nomor antrean => generate nomor antrean berdasarkan tanggal,.
kode_booking => ,
GENERATE KODE BOKING LAGI DI kode_booking.

GENERATE KODE BOOKING
2023 01 13 => tahun - bulan - tanggal
nomor 0001



SELECT no_antrean, no_antrean_angka, tanggal, kode_booking, SUM(dilayani="false") AS sisaantrean, SUM(dilayani="true") AS antreanpanggil, 
COUNT(dilayani) AS totalantrean FROM posfar.antrean_resep WHERE  tanggal="2023-01-13"


