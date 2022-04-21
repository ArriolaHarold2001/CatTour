import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export class IImg {
  constructor(public url: string) {}
}

@Component({
  selector: 'app-cat-of-day',
  templateUrl: './CatofDay.component.html',
  styleUrls: ['./CatOfDay.component.scss'],
})
export class CatOfDay implements OnInit {
  catImg: IImg[] = [];
  constructor(private httpClient: HttpClient) {}

  ngOnInit(): void {
    this.getCatImg();
  }

  getCatImg() {
    this.httpClient.get<any>('http://127.0.0.1:8000').subscribe((response) => {
      console.log(response);
      this.catImg = response;
    });
  }
}
