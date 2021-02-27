import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Advertisement } from 'src/app/models/ad';
import { AdService } from 'src/app/services/ad/ad.service';

@Component({
  selector: 'app-ad-page',
  templateUrl: './ad-page.component.html',
  styleUrls: ['./ad-page.component.scss']
})
export class AdPageComponent implements OnInit {

  constructor(
    private adService: AdService,
    private route: ActivatedRoute,
    private location: Location
  ) { }

  fetchPending = true;
  ad: Advertisement;

  ngOnInit(): void {
    this.adService.getOne(+this.route.snapshot.params.id).subscribe(
      (ad: Advertisement) => {
        this.fetchPending = false;
        if (ad){
          this.ad = ad;
        }
        else{
          this.location.back();
        }
      }
    );
  }

}
