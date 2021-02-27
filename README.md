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
Arhitektura sistema bazirana je mikroservisima. Svaki mikroservis poseduje zasebnu bazu podataka (konkretno PostgreSQL) i u njoj cuva podatke kojima samo on upravlja. Sva komunikacija izmedju klijenta i servisa odvija se preko REST API-a. Mikroservis za autentifikaciju (Users Microservice) definise SECRET_KEY koji ostali servisi koriste za dekodovanje JWT tokena i dobavljanje podataka o prijavljenom korisniku.

<h2>Pregled mikroservisa</h2>
<h6>Users Microservice</h6>
Django REST aplikacija koja omogucava prijavu, registraciju (uz verifikaciju email-a), izmenu profila korisnika i koriscenje ugradjenog Django admin sistema za administraciju korisnika (kreiranje, brisanje, izmena i pregled). Servis prilikom prijave generise JWT token koji ce korisnik koristiti za autentifikaciju na svim servisima i definise SECRET_KEY koji ce ostali servisi koristiti za dekodovanje JWT tokena. <b>Port mikroservisa je 8000.</b> Mikroservis se pokrece komandom: <b>python manage.py runserver </b>. 
<h6>Advertisements Microservice</h6>
Golang REST aplikacija koja omogucava authentifikovanim korisnicima kreiranje, izmenu, brisanje, pregled, paginaciju i pretragu mednih proizvoda. <b>Podaci kojima je opisana reklama je: </b>datum objave, ime proizvoda koji se reklamira, kategorija proizvoda, cena proizvoda, opis proizvoda i skup slika proizvoda. <b>Port mikroservisa je 8001.</b> Mikroservis se pokrece komandom <b>go run main.go</b>. 
<h6>Events Microservice</h6>
Golang REST aplikacija koja omogucava autentifikovanim korisnicima kreiranje, izmenu, brisanje, pregled i paginaciju dogadjaja na kojima se prezentuje neki medni proizvod. <b>Podaci kojima je opisan dogadjaj je: </b>datum objave, ime dogadjaja, kategorija dogadjaja (sajam, manifestacija...), period i mesto odrzavanja dogadjaja, opis dogadjaja i skup slika. <b>Port mikroservisa je 8002. </b>Mikroservis se pokrece komandom <b>go run main.go</b>.
<h6>Comments Microservice</h6>
Django REST aplikacija koja omogucava komentarisanje reklamiranih proizvoda, pregled i paginaciju komantara i podkomentara i like/dislike komentara. <b>Port mikroservisa je 8003. </b> Mikroservis se pokrece komandom <b>python manage.py runserver 8003</b>. 

<h2>Uputstvo za pokretanje</h2>
1. Koristeci komandu <b>python -m venv venv</b> (ili python3 -m venv venv ako su na racunaru instalirani i pajton2 i pajton3) kreirati virtuelno okruzenje
2. U virtuelno okruzenje instalirati sve biblioteke navedene u <b>requirements.txt</b> fajlu
3. Aktivirati virtuelno okruzenje, pozicionirati se u <b>user_service</b> i pokrenuti komandu <b>python manage.py runserver</b>
4. Aktivirati virtuelno okruzenje, pozitionirati se u <b>comment_service</b> i pokrenuti komandu <b>python manage.py runserver 8003</b>
5. Pokrenuti komande za preuzimanje neophodnih Goland biblioteka:
<ol>
  <li>go get -u -v github.com/dgrijalva/jwt-go</li>
  <li>go get -u -v github.com/gorilla/mux</li>
  <li>go get -u -v github.com/jinzhu/gorm</li>
  <li>go get -u -v github.com/lib/pq</li>
  <li>go get -u -v github.com/rs/cors</li>
</ol>
6. Pozicionirati se u <b>ad_service</b> i pokrenuti komandu <b>go run main.go</b>
7. Pozicionirati se u <b>event_service</b> i pokrenuti komandru <b>go run main.go</b>
8. Pozicionirati se u <b>angular-client</b> i pokrenuti komande <b>npm install</b> i <b>ng serve</b>
9. U URL browsera uneti putanju <b>localhost:4200</b> ako zelite da koristiti Angular klijenta
10. TODO: Pharo klijent

<br><br><br>
5. Client<br>
Sistem ce imati jednu frontend aplikaciju koja ce pozivate metode sva cetri navedena servisa i u pitanju ce biti Angular aplikacija. 
