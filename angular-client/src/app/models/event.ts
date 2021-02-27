import { Image } from './image';

export interface Event{
    ID: number;
    CreatedOn: string;
    UserId: number;
    Email: string;
    ProductId: number;
    Name: string;
    Category: string;
    From: string;
    To: string;
    Place: string;
    Description: string;
    Images: Image[];
}
