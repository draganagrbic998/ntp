import { HttpHeaders, HttpResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FIRST_PAGE_HEADER, LAST_PAGE_HEADER } from 'src/app/constants/pagination';
import { Pagination } from 'src/app/models/pagination';
import { AdService } from 'src/app/services/ad/ad.service';
import { Advertisement } from 'src/app/models/ad';

@Component({
  selector: 'app-ad-list',
  templateUrl: './ad-list.component.html',
  styleUrls: ['./ad-list.component.scss']
})
export class AdListComponent implements OnInit {

  constructor(
    private adService: AdService
  ) { }

  ads: Advertisement[] = [];
  fetchPending = true;
  pagination: Pagination = {
    pageNumber: 0,
    firstPage: true,
    lastPage: true
  };
  search = '';

  changePage(value: number): void{
    this.pagination.pageNumber += value;
    this.fetchAds();
  }

  fetchAds(): void{
    this.fetchPending = true;
    this.adService.getAll(this.pagination.pageNumber, this.search).subscribe(
      (data: HttpResponse<Advertisement[]>) => {
        this.fetchPending = false;
        if (data){
          this.ads = data.body;
          const headers: HttpHeaders = data.headers;
          this.pagination.firstPage = headers.get(FIRST_PAGE_HEADER) === 'false' ? false : true;
          this.pagination.lastPage = headers.get(LAST_PAGE_HEADER) === 'false' ? false : true;
        }
        else{
          this.ads = [];
          this.pagination.firstPage = true;
          this.pagination.lastPage = true;
        }
      }
    );
  }

  ngOnInit(): void {
    this.changePage(0);
    this.adService.refreshData$.subscribe(() => {
      this.changePage(0);
    });
    this.adService.searchData$.subscribe((search: string) => {
      this.search = search;
      this.pagination.pageNumber = 0;
      this.changePage(0);
    });
  }

}
