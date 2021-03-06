import { Component, OnInit } from '@angular/core';

import { Blueprint, BlueprintsService } from '../../services/blueprints/blueprints';
import { IterateMapPipe, KeysPipe } from '../../pipes/iterate';

@Component({
  selector: 'my-dashboard',
   pipes: [IterateMapPipe, KeysPipe],
  providers: [BlueprintsService],
  styleUrls: ['app/components/dashboard/dashboard.css'],
  templateUrl: 'app/components/dashboard/dashboard.html',
})

export class DashboardComponent implements OnInit {

  blueprints: Blueprint[];
  error: any;

  constructor(
    private blueprintsService: BlueprintsService
  ) { }

  ngOnInit() {
    this.getBlueprints();
  }

  getBlueprints() {
    this.blueprintsService.getBlueprints()
      .subscribe(blueprints => this.blueprints = blueprints)
  }

}
