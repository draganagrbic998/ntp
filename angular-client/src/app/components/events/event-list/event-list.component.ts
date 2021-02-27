import { HttpHeaders, HttpResponse } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { FIRST_PAGE_HEADER, LAST_PAGE_HEADER } from 'src/app/constants/pagination';
import { Event } from 'src/app/models/event';
import { Pagination } from 'src/app/models/pagination';
import { EventService } from 'src/app/services/event/event.service';

@Component({
  selector: 'app-event-list',
  templateUrl: './event-list.component.html',
  styleUrls: ['./event-list.component.scss']
})
export class EventListComponent implements OnInit {

  constructor(
    private eventService: EventService
  ) { }

  @Input() productId: number;
  events: Event[] = [];
  fetchPending = true;
  pagination: Pagination = {
    pageNumber: 0,
    firstPage: true,
    lastPage: true
  };

  changePage(value: number): void{
    this.pagination.pageNumber += value;
    this.fetchEvents();
  }

  fetchEvents(): void{
    this.fetchPending = true;
    this.eventService.getAll(this.pagination.pageNumber, this.productId).subscribe(
      (data: HttpResponse<Event[]>) => {
        this.fetchPending = false;
        if (data){
          this.events = data.body;
          const headers: HttpHeaders = data.headers;
          this.pagination.firstPage = headers.get(FIRST_PAGE_HEADER) === 'false' ? false : true;
          this.pagination.lastPage = headers.get(LAST_PAGE_HEADER) === 'false' ? false : true;
        }
        else{
          this.events = [];
          this.pagination.firstPage = true;
          this.pagination.lastPage = true;
        }
      }
    );
  }

  ngOnInit(): void {
    this.changePage(0);
    this.eventService.refreshData$.subscribe(() => {
      this.changePage(0);
    });
  }

}
