import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Image } from 'src/app/models/image';

@Component({
  selector: 'app-image-input',
  templateUrl: './image-input.component.html',
  styleUrls: ['./image-input.component.scss']
})
export class ImageInputComponent implements OnInit {

  constructor() { }

  @Input() image: Image;
  @Output() deleted: EventEmitter<null> = new EventEmitter();

  ngOnInit(): void {
  }

}
