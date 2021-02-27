import { Injectable } from '@angular/core';
import { User } from 'src/app/models/user';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor() { }

  private readonly STORAGE_KEY = 'auth';

  saveUser(user: User): void{
    localStorage.setItem(this.STORAGE_KEY, JSON.stringify(user));
  }

  deleteUser(): void{
    localStorage.removeItem(this.STORAGE_KEY);
  }

  getUser(): User{
    return JSON.parse(localStorage.getItem(this.STORAGE_KEY));
  }

}
