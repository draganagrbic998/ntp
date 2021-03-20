import { Image } from './image';

export interface Advertisement{
    ID: number;
    CreatedOn: string;
    UserID: number;
    Name: string;
    Category: string;
    Price: number;
    Description: string;
    Images: Image[];
}
