import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MultipleAddComponent } from './multiple-add.component';

describe('MultipleAddComponent', () => {
  let component: MultipleAddComponent;
  let fixture: ComponentFixture<MultipleAddComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MultipleAddComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MultipleAddComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
