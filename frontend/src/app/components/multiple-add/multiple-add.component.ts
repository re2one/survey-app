import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {MuchoService} from '../../services/mucho.service';
import {HttpResponse} from '@angular/common/http';
import {AnswerResponse} from '../../models/mucho';

@Component({
  selector: 'app-multiple-add',
  templateUrl: './multiple-add.component.html',
  styleUrls: ['./multiple-add.component.css']
})
export class MultipleAddComponent implements OnInit {
  questionId: string;
  surveyId: string;
  constructor(
    public router: Router,
    private multipleService: MuchoService,
    private activatedRoute: ActivatedRoute,
  ) { }

  ngOnInit(): void {
    this.activatedRoute.paramMap.subscribe(params => {
      this.questionId = params.get('questionId');
      this.surveyId = params.get('surveyId');
    });
  }
  onAnswerSubmit(surveyData): void{
    this.multipleService.postAnswer(
      surveyData.text,
      this.questionId,
    ).subscribe((response: HttpResponse<AnswerResponse>) => {
      console.log(response);
      if (response.status === 200) {
        this.router.navigate(['/questions/edit/multiplechoice', this.questionId, this.surveyId]);
      }
    });
  }
}
