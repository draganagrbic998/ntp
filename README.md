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
U pitanju je mikroservisna arhitektura (nismo na fakultetu radili takav projekat, pa mi se to bas svidelo :D). Svaki mikroservis ima svoju posebnu bazu kojom samo on upravlja. Sve baze u sistemu ce biti Postgre baze i koristicu ORM mapiranje (i kod Python aplikacija i kod Golang aplikacija). Sva komunikacija izmedju klijenta i servisa odvijace se preko REST API-a. Stvari koje planiram da implementiram su:
1. Users Service<br>
Djnago REST aplikacija koja omogucava prijavu korisnika, registraciju (uz verifikaciju email-a), izmenu profila korisnika i koriscenje ugradjenog Django admin sistema za administraciju korisnika (kreiranje, brisanje, izmena i pregled). Servis prilikom prijave generise JWT token koji ce korisnik koristiti za autentifikaciju na ostalim servisima i definise SECRET_KEY koji ce ostali servisi koristiti za dekodovanje JWT tokena. <b>Port mikroservisa je 8000.</b> Mikroservis se pokrece komandom: <b>python manage.py runserver </b>. 
2. Advertisements Service<br>
Golang REST aplikacija koja omogucava authentifikovanim korisnicima kreiranje, izmenu, brisanje, pregled, paginaciju i pretragu mednih proizvoda. Servis koristi SECRET_KEY users mikroservisa da dekoduje JWT token i dobavi neophodne podatke o korisniku. <b>Podaci kojima je opisana reklama je: </b>datum objave, ime proizvoda koji se reklamira, kategorija proizvoda, cena proizvoda, opis proizvoda i skup slika proizvoda. <b>Port mikroservisa je 8001.</b> Mikroservis se pokrece komandom <b>go run main.go</b>. 
3. Comment Service<br>
Ovo ce biti Django Rest aplikacija koja ce omogucavati gostima da komentarisu medne proizvode. Komentari ce imati podkomentare, lajkove i dislajkove. Gosti mogu da postave komentare, mogu da ih obrisu, mogu da menjaju svoje komentare.
4. Event Service<br>
Ovo ce biti Golang aplikacija koja ce nuditi REST API koji ce pcelarima omoguciti da kreiraju dogadjaje na kojima ce prezentovati svoje ukusne proizvode (sajam, manifestacija itd.). Dogadjaj planiram da opisem jednostavnim atributima (datum pocetka/kraja, naziv, opis, mesto itd.) i skupom slika. Pcelari ce moci da kreiraju novi dogadjaj, da ga izmene, obrisu, pregledaju i filtriraju. Gosti ce moci da pregledaju dogadjaje i pretrazuju.
5. Client<br>
Sistem ce imati jednu frontend aplikaciju koja ce pozivate metode sva cetri navedena servisa i u pitanju ce biti Angular aplikacija. 
Neka odstupanja od gore navedenih funkcionalnosti mozda uvedem kako budem implementirala projekat, al sustina ce ostati ista.
<br>
<b>Napomena:</b><br>
Znam da u obavestenju pise da je obavezno koristiti Pharo i Golang jezike u projektu, a da ja Python opcion, ali nisam mogla da smislim gde bih mogla da koristim Pharo u ovakvom projektu, osim eventualno da vizualizujem neku zavisnost (npr. koliko se koja kategorija proizvoda reklamira ili tako nesto). Da li je obavezno da koristim Pharo u ovom projektu onda? Ako jeste, da li je u redu ga iskoristim da npr. vizualizujem koliko se koje kategorije reklamiraju ili nesto slicno? 
