import { Image } from './image';

export interface Event{
    ID: number;
    CreatedOn: string;
    UserID: number;
    ProductID: string;
    Name: string;
    Category: string;
    From: string;
    To: string;
    Place: string;
    Description: string;
    Images: Image[];
}
