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
Golang REST aplikacija koja omogucava authentifikovanim korisnicima kreiranje, izmenu, brisanje, pregled, paginaciju i pretragu reklama mednih proizvoda. <b>Podaci kojima je opisana reklama je: </b>datum objave, ime proizvoda koji se reklamira, kategorija proizvoda, cena proizvoda, opis proizvoda i skup slika proizvoda. <b>Port mikroservisa je 8001.</b>
<h6>Events Microservice</h6>
Golang REST aplikacija koja omogucava autentifikovanim korisnicima kreiranje, izmenu, brisanje, pregled i paginaciju dogadjaja na kojima se prezentuje neki medni proizvod. <b>Podaci kojima je opisan dogadjaj je: </b>datum objave, ime dogadjaja, kategorija dogadjaja (sajam, manifestacija...), period i mesto odrzavanja dogadjaja, opis dogadjaja i skup slika. <b>Port mikroservisa je 8002.</b>
<h6>Comments Microservice</h6>
Django REST aplikacija koja omogucava komentarisanje reklamiranih proizvoda, pregled i paginaciju komantara i podkomentara i like/dislike komentara. <b>Port mikroservisa je 8003.</b>

<br><h2>Klijenti sistema</h2>
<h6>Angular klijent</h6>
Glavni klijent implementiran u Angular jeziku koji omogucava koriscenje glavnih funkcionalnosti sistema. Lokacija klijenta je <b>localhost:4200</b>.
<h6>Pharo klijent</h6>
Klijent koji omogucava uvid u analitiku sistema. Kada administrator unutar Pharo okruzenja registruje klasu sa metodom koja se nalazi na putanji <b>pharo-client/graphics.txt</b>, dobija mogucnost slanja poruka okruzenju koje ce mu omoguciti graficki prikaz broja reklama, dogadjaja, komentara, lajkova i dislajkova u odabranom vremenskom intervalu. Format poruka je <b>NAZIV_KLASE entity ENTITET start START end END shape SHAPE</b>, gde su elementi redom:
<ul>
  <li>
    NAZIV_KLASE predstavlja naziv pod kojim je korisnik/admin registrovao klasu koja ce implementirati metodu za graficke prikaze.
  </li>
  <li>
    ENTITET predstavlja izbor entiteta ciji graficki prikaz se zeli prikazati. Validne vrednosti su "ads", "events", "comments", "likes" i "dislikes".
  </li>
  <li>
    START predstavlja pocetnu godinu od koje se analizira statistika odabranog entiteta. Godina mora biti cetvorocifren broj.
  </li>
  <li>
    END predstavlja krajnju godinu do koje se analizira statistika odabranog entiteta. Godina mora biti cetvorocifren broj, mora biti veca od pocetne godine i ne sme biti veca od tekuce godine. 
  </li>
  <li>
    SHAPE omogucava korisniku odabir prikaza grafika u vidu plot-a (ako se unese "dots") i bar-ova (ako se unese "bars").
  </li>
</ul>
<h6>Django admin aplikacija</h6>
Users mikroservis pruza koriscenje ugradjene Django admin aplikacije koja omogucava administraciju korisnika - kreiranje, izmena, brisanje i pregled. Lokacija klijenta je <b>localhost:8000/admin</b>.

<br><h2>Uputstvo za pokretanje</h2>
<ol>
  <li>
    Lolalno kreirati Postgre baze koje ce se zvati: <b>users</b>, <b>ads</b>, <b>events</b> i <b>comments</b>.
  </li>
  <li>
    Koristeci komandu <b>python -m venv venv</b> (ili python3 -m venv venv ako su na racunaru instalirani i pajton2 i pajton3) kreirati virtuelno okruzenje.
  </li>
  <li>
    U virtuelno okruzenje instalirati sve biblioteke navedene u <b>requirements.txt</b> fajlu.
  </li>
  <li>
    Aktivirati virtuelno okruzenje, pozicionirati se u <b>user_service</b> i pokrenuti komande <b>python manage.py migrate</b> i <b>python manage.py runserver</b>.
  </li>
  <li>
    Aktivirati virtuelno okruzenje, pozitionirati se u <b>comment_service</b> i pokrenuti komande <b>python manage.py migrate</b> i <b>python manage.py runserver 8003</b>.
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
    Pozicionirati se u <b>ad_service</b> i pokrenuti komandu <b>go run .</b>.
  </li>
  <li>
    Pozicionirati se u <b>event_service</b> i pokrenuti komandu <b>go run .</b>.
  </li>
  <li>
    Pozicionirati se u <b>angular-client</b> i pokrenuti komande <b>npm install</b> i <b>ng serve</b>.
  </li>
  <li>
    Unutar Pharo okruzenja instalirati biblioteke <b>Roassal</b> i <b>NeoJSON</b>. 
  </li>
  <li>
    Unutar Pharo okruzenja registrovati/kreiratu klasu u proizvoljnom paketu i sa proizvolnjim nazivom i unutar <b>Class opsega</b> novokreirane klase registrovati metodu koja se nalazi na putanji <b>pharo-client/graphics.txt</b>.
  </li>
  <li>
    U URL browsera uneti putanju <b>localhost:4200</b> ukoliko zelite da koristiti Angular klijenta.
  </li>
  <li>
    Otvoriti Playground Pharo okruzenja i slati prethodno opisane poruke ukoliko zelite da koristiti Pharo klijenta.
  </li>
  <li>
    Na URL-u <b>localhost:8000/admin</b> mozete vrsiti administraciju korisnika.
  </li>
</ol>

<br><h2>Pokretanje testova</h2>
<h6>Pokretanje jedinicnih testova mikroservisa napisanih u Python-u</h6>
Pozicionirati se u direktorijum mikroservisa i pokrenuti komandu <b>python manage.py test</b>.
<h6>Pokretanje jedinicnih testova mikroservisa napisanih u Golang-u</h6>
Pozicionirati se u direktorijum mikroservisa i pokrenuti komandu <b>go test -v</b>.
