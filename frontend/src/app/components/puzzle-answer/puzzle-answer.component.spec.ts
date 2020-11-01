import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PuzzleAnswerComponent } from './puzzle-answer.component';

describe('PuzzleAnswerComponent', () => {
  let component: PuzzleAnswerComponent;
  let fixture: ComponentFixture<PuzzleAnswerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PuzzleAnswerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PuzzleAnswerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
