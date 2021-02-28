<h1>Projekat iz predmeta Napredne tehnike programiranja</h1>

<div align="center">
  <img src="https://github.com/draganagrbic998/ntp/blob/main/pcelica.jpg" alt="drawing" width="200" height="200"/>
</div>

<h2>Opis problema</h2>
Neophodno je implementirati sistem koji ce omoguciti precalirima i ljubiteljima mednih proizvoda da reklamiraju svoja dobra, pretrazuju i komentarisu tudja i obavestavaju ostale korisnike o dogadjajima na kojima ce prezentovati svoje medne proizvode.

<br><h2>Arhitektura sistema</h2>
Arhitektura sistema bazirana je mikroservisima. Svaki mikroservis poseduje zasebnu bazu podataka (konkretno PostgreSQL) i u njoj cuva podatke kojima samo on upravlja. Sva komunikacija izmedju klijenta i servisa odvija se preko REST API-a. Mikroservis za autentifikaciju (Users Microservice) definise SECRET_KEY koji ostali servisi koriste za dekodovanje JWT tokena i dobavljanje podataka o prijavljenom korisniku.

<br><h2>Pregled mikroservisa</h2>
<h6>Users Microservice</h6>
Django REST aplikacija koja omogucava prijavu, registraciju (uz verifikaciju email-a), izmenu profila korisnika i koriscenje ugradjenog Django admin sistema za administraciju korisnika (kreiranje, brisanje, izmena i pregled). Servis prilikom prijave generise JWT token koji ce korisnik koristiti za autentifikaciju na svim servisima i definise SECRET_KEY koji ce ostali servisi koristiti za dekodovanje JWT tokena. <b>Port mikroservisa je 8000.</b> 
<h6>Advertisements Microservice</h6>
Golang REST aplikacija koja omogucava authentifikovanim korisnicima kreiranje, izmenu, brisanje, pregled, paginaciju i pretragu mednih proizvoda. <b>Podaci kojima je opisana reklama je: </b>datum objave, ime proizvoda koji se reklamira, kategorija proizvoda, cena proizvoda, opis proizvoda i skup slika proizvoda. <b>Port mikroservisa je 8001.</b>
<h6>Events Microservice</h6>
Golang REST aplikacija koja omogucava autentifikovanim korisnicima kreiranje, izmenu, brisanje, pregled i paginaciju dogadjaja na kojima se prezentuje neki medni proizvod. <b>Podaci kojima je opisan dogadjaj je: </b>datum objave, ime dogadjaja, kategorija dogadjaja (sajam, manifestacija...), period i mesto odrzavanja dogadjaja, opis dogadjaja i skup slika. <b>Port mikroservisa je 8002.</b>
<h6>Comments Microservice</h6>
Django REST aplikacija koja omogucava komentarisanje reklamiranih proizvoda, pregled i paginaciju komantara i podkomentara i like/dislike komentara. <b>Port mikroservisa je 8003.</b>

<br><h2>Klijenti sistema</h2>
U sistemu implementirana su tri klijenta:
<h6>Angular klijent</h6>
Glavni klijent implementiran u Angular jeziku koji omogucava koriscenje glavnih funkcionalnosti sistema.
<h6>Pharo okruzenje</h6>
TODO :D
<h6>Django admin aplikacija</h6>
Users mikroservis pruza koriscenje ugradjene Django admin aplikacije koja omogucava administraciju korisnika - kreiranje, izmena, brisanje i pregled.

<br><h2>Uputstvo za pokretanje</h2>
<ol>
  <li>
    Koristeci komandu <b>python -m venv venv</b> (ili python3 -m venv venv ako su na racunaru instalirani i pajton2 i pajton3) kreirati virtuelno okruzenje
  </li>
  <li>
    U virtuelno okruzenje instalirati sve biblioteke navedene u <b>requirements.txt</b> fajlu
  </li>
  <li>
    Aktivirati virtuelno okruzenje, pozicionirati se u <b>user_service</b> i pokrenuti komandu <b>python manage.py runserver</b>
  </li>
  <li>
    Aktivirati virtuelno okruzenje, pozitionirati se u <b>comment_service</b> i pokrenuti komandu <b>python manage.py runserver 8003</b>
  </li>
  <li>
    Pokrenuti komande za preuzimanje neophodnih Golang biblioteka:
    <ul>
      <li>go get -u -v github.com/dgrijalva/jwt-go</li>
      <li>go get -u -v github.com/gorilla/mux</li>
      <li>go get -u -v github.com/jinzhu/gorm</li>
      <li>go get -u -v github.com/lib/pq</li>
      <li>go get -u -v github.com/rs/cors</li>
    </ul>
  </li>
  <li>
    Pozicionirati se u <b>ad_service</b> i pokrenuti komandu <b>go run main.go</b>
  </li>
  <li>
    Pozicionirati se u <b>event_service</b> i pokrenuti komandru <b>go run main.go</b>
  </li>
  <li>
    Pozicionirati se u <b>angular-client</b> i pokrenuti komande <b>npm install</b> i <b>ng serve</b>
  </li>
  <li>
    U URL browsera uneti putanju <b>localhost:4200</b> ako zelite da koristiti Angular klijenta
  </li>
</ol>
