import { Component, Input, OnInit } from '@angular/core';
import { Image } from 'src/app/models/image';
import { ImageService } from 'src/app/services/image/image.service';

@Component({
  selector: 'app-images-input',
  templateUrl: './images-input.component.html',
  styleUrls: ['./images-input.component.scss']
})
export class ImagesInputComponent implements OnInit {

  constructor(
    private imageService: ImageService
  ) { }

  @Input() images: Image[] = [];

  addImage(upload: Blob): void{
    // tslint:disable-next-line: deprecation
    this.imageService.getBase64(upload).subscribe(
      (Path: string) => {
        this.images.push({Path} as Image);
      }
    );
  }

  ngOnInit(): void {
  }

}
