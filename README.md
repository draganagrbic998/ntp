<div align="center">
  <img src="https://github.com/draganagrbic998/ntp/blob/main/pcelica.jpg" alt="drawing" width="200" height="200"/>
</div>

# Projekat iz predmeta Napredne tehnike programiranja
Tema projekta (samostalno definisani projekat):<br>

Implementacija sistema koja ce omogucavati pcelarima da reklamiraju svoje proizvoda i ljubiteljima mednih proizvoda da pretrazuju te proizvode i komentarisu ih.
Korisnici sistema bice:
1. pcelari - kreiraju svoje proizvode i dogadjaje na kojima ce ih prezentovati
2. ljubitelji mednih proizvoda (u daljem tekstu zvacu ih GOST) - komentarisu medne proizvode, pretrazuju i pregledaju ih
Izgled arhitekture bio bi sledeci:
<br><br><br>

![alt text](https://github.com/draganagrbic998/ntp/blob/main/ntp_diagram.png)
<br><br>

<h2>Arhitektura sistema</h2>
Arhitektura sistema bazirana je mikroservisima. Svaki mikroservis poseduje zasebnu bazu podataka (konkretno PostgreSQL) i u njoj cuva podatke kojima samo ona upravlja. Sva komunikacija izmedju klijenta i servisa odvija se preko REST API-a. Mikroservis za autentifikaciju (Users Microservice) definise SECRET_KEY koji ostali servisi koriste za dekodovanje JWT tokena i dobavljanje podataka o prijavljenom korisniku.

<h2>Pregled mikroservisa</h2>
<h6>Users Microservice</h6>
Djnago REST aplikacija koja omogucava prijavu, registraciju (uz verifikaciju email-a), izmenu profila korisnika i koriscenje ugradjenog Django admin sistema za administraciju korisnika (kreiranje, brisanje, izmena i pregled). Servis prilikom prijave generise JWT token koji ce korisnik koristiti za autentifikaciju na svim servisima i definise SECRET_KEY koji ce ostali servisi koristiti za dekodovanje JWT tokena. <b>Port mikroservisa je 8000.</b> Mikroservis se pokrece komandom: <b>python manage.py runserver </b>. 
<h6>Advertisements Microservice</h6>
Golang REST aplikacija koja omogucava authentifikovanim korisnicima kreiranje, izmenu, brisanje, pregled, paginaciju i pretragu mednih proizvoda. <b>Podaci kojima je opisana reklama je: </b>datum objave, ime proizvoda koji se reklamira, kategorija proizvoda, cena proizvoda, opis proizvoda i skup slika proizvoda. <b>Port mikroservisa je 8001.</b> Mikroservis se pokrece komandom <b>go run main.go</b>. 
<h6>Events Microservice</h6>
Golang REST aplikacija koja omogucava autentifikovanim korisnicima kreiranje, izmenu, brisanje, pregled i paginaciju dogadjaja na kojima se prezentuje neki medni proizvod. <b>Podaci kojima je opisan dogadjaj je: </b>datum objave, ime dogadjaja, kategorija dogadjaja (sajam, manifestacija...), period i mesto odrzavanja dogadjaja, opis dogadjaja i skup slika. <b>Port mikroservisa je 8002. </b>Mikroservis se pokrece komandom <b>go run main.go</b>.
<h6>Comments Microservice</h6>
Django REST aplikacija koja omogucava komentarisanje reklamiranih proizvoda, pregled i paginaciju komantara i podkomentara i like/dislike komentara. <b>Port mikroservisa je 8003. </b> Mikroservis se pokrece komandom <b>python manage.py runserver 8003</b>. 

<br><br><br>
5. Client<br>
Sistem ce imati jednu frontend aplikaciju koja ce pozivate metode sva cetri navedena servisa i u pitanju ce biti Angular aplikacija. 
