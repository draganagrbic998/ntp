import { HttpHeaders, HttpResponse } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { FIRST_PAGE_HEADER, LAST_PAGE_HEADER } from 'src/app/constants/pagination';
import { Comment } from 'src/app/models/comment';
import { Pagination } from 'src/app/models/pagination';
import { CommentService } from 'src/app/services/comment/comment.service';

@Component({
  selector: 'app-comment-list',
  templateUrl: './comment-list.component.html',
  styleUrls: ['./comment-list.component.scss']
})
export class CommentListComponent implements OnInit {

  constructor(
    private commentService: CommentService
  ) { }

  @Input() productId: number;
  comments: Comment[] = [];
  fetchPending = true;
  pagination: Pagination = {
    pageNumber: 0,
    firstPage: true,
    lastPage: true
  };

  changePage(value: number): void{
    this.pagination.pageNumber += value;
    this.fetchComments();
  }

  fetchComments(): void{
    this.fetchPending = true;
    this.commentService.getAll(this.pagination.pageNumber, this.productId).subscribe(
      (data: HttpResponse<Comment[]>) => {
        this.fetchPending = false;
        if (data){
          this.comments = data.body;
          const headers: HttpHeaders = data.headers;
          this.pagination.firstPage = headers.get(FIRST_PAGE_HEADER) === 'false' ? false : true;
          this.pagination.lastPage = headers.get(LAST_PAGE_HEADER) === 'false' ? false : true;
        }
        else{
          this.comments = [];
          this.pagination.firstPage = true;
          this.pagination.lastPage = true;
        }
      }
    );
  }

  ngOnInit(): void {
    this.changePage(0);
    this.commentService.refreshData$.subscribe(() => {
      this.changePage(0);
    });
  }

}
