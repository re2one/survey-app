import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {MuchoService} from '../../services/mucho.service';
import {HttpResponse} from '@angular/common/http';
import {SurveyResponse} from '../../models/survey';

@Component({
  selector: 'app-multiple-edit',
  templateUrl: './multiple-edit.component.html',
  styleUrls: ['./multiple-edit.component.css']
})
export class MultipleEditComponent implements OnInit {

  answerId: string;
  surveyId: string;
  constructor(
    public router: Router,
    private multipleService: MuchoService,
    private activatedRoute: ActivatedRoute,
  ) { }

  ngOnInit(): void {
    this.activatedRoute.paramMap.subscribe(params => {
      this.answerId = params.get('answerId');
      this.surveyId = params.get('surveyId');
    });
  }
  onAnswerSubmit(answerData): void{
    this.multipleService.putAnswer(
      parseInt(this.answerId, 10),
      answerData.questionid,
      answerData.text,
      answerData.nextQuestion,
      answerData.secondToNext,
      answerData.typeOfNextQuestion,
    ).subscribe((response: HttpResponse<SurveyResponse>) => {
      if (response.status === 200) {
        this.router.navigate(['/questions/edit/multiplechoice', answerData.questionId, this.surveyId]);
      }
    });
  }

}
