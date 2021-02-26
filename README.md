# Projekat iz predmeta Napredne tehnike programiranja

Tema projekta (samostalno definisani projekat):<br>
Implementacija sistema koja ce omogucavati pcelarima da reklamiraju svoje proizvoda i ljubiteljima mednih proizvoda da pretrazuju te proizvode i komentarisu ih.
Korisnici sistema bice:
1. pcelari - kreiraju svoje proizvode i dogadjaje na kojima ce ih prezentovati
2. ljubitelji mednih proizvoda (u daljem tekstu zvacu ih GOST) - komentarisu medne proizvode, pretrazuju i pregledaju ih
Izgled arhitekture bio bi sledeci:

![alt text](https://github.com/draganagrbic998/ntp/blob/main/ntp_diagram.png)

U pitanju je mikroservisna arhitektura (nismo na fakultetu radili takav projekat, pa mi se to bas svidelo :D). Svaki mikroservis ima svoju posebnu bazu kojom samo on upravlja. Sve baze u sistemu ce biti Postgre baze i koristicu ORM mapiranje (i kod Python aplikacija i kod Golang aplikacija). Sva komunikacija izmedju klijenta i servisa odvijace se preko REST API-a. Stvari koje planiram da implementiram su:
1. User Service
Ovo ce biti Django aplikacija (preciznije koristicu Django Rest Framework) i ona ce implementirati prijavu (zajedno sa slanjem JWT tokena kojim ce se korisnik autentifikovati ostatku sistema), registraciju korisnika (zajedno sa verifikacijom email-a), pregled profila od strane korisnika i izmena profila.
2. Product Service
Ovo ce biti Golang aplikacija koja se nuditi REST API koji ce pcelarima omoguciti da kreiraju proizvode koje prave kod sebe. Proizvod planiram da opisem jednostanim atributima (naziv, opis, kategorija, tip, cena, itd.) i skupom slika. Pcelar ce moci da kreira svoje proizvode, menja, brise, pregleda, pretrazuje. Gosti mogu da pregledaju proizvode, filtriraju po razlicitim kriterijumima itd. 
3. Comment Service
Ovo ce biti Django Rest aplikacija koja ce omogucavati gostima da komentarisu medne proizvode. Komentari ce imati podkomentare, lajkove i dislajkove. Gosti mogu da postave komentare, mogu da ih obrisu, mogu da menjaju svoje komentare. 
4. Event Service
Ovo ce biti Golang aplikacija koja ce nuditi REST API koji ce pcelarima omoguciti da kreiraju dogadjaje na kojima ce prezentovati svoje ukusne proizvode (sajam, manifestacija itd.). Dogadjaj planiram da opisem jednostavnim atributima (datum pocetka/kraja, naziv, opis, mesto itd.) i skupom slika. Pcelari ce moci da kreiraju novi dogadjaj, da ga izmene, obrisu, pregledaju i filtriraju. Gosti ce moci da pregledaju dogadjaje i pretrazuju. 
5. Client
Sistem ce imati jednu frontend aplikaciju koja ce pozivate metode sva cetri navedena servisa i u pitanju ce biti Angular aplikacija. 

Neka odstupanja od gore navedenih funkcionalnosti mozda uvedem kako budem implementirala projekat, al sustina ce ostati ista.

