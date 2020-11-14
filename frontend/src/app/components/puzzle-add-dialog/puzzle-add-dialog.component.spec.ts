import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PuzzleAddDialogComponent } from './puzzle-add-dialog.component';

describe('PuzzleAddDialogComponent', () => {
  let component: PuzzleAddDialogComponent;
  let fixture: ComponentFixture<PuzzleAddDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PuzzleAddDialogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PuzzleAddDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
