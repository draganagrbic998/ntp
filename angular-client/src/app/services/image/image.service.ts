import { Injectable } from '@angular/core';
import { from, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ImageService {

  constructor() { }

  getBase64(upload: Blob): Observable<string>{
    return from(new Promise<string>((resolve, reject) => {
      const reader: FileReader = new FileReader();
      reader.onload = () => resolve(reader.result as string);
      reader.onerror = () => reject();
      reader.readAsDataURL(upload);
    }));
  }

}
