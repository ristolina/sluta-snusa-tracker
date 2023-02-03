from django.shortcuts import render
from datetime import datetime
from django.http import HttpResponseRedirect
import requests
import logging
import json
import os
from datetime import datetime
from dateutil.relativedelta import relativedelta
from django.contrib.auth.decorators import login_required

# Create your views here.

@login_required
def homeView(request):
    username = request.user.get_username()
    requestUrl = "http://backend:8080/slutasnusa/" + username
    response = requests.get(requestUrl)
    data = json.loads(response.text)
    data = data['message']
    if not isinstance(data, str):
        quitDateTime = datetime.strptime(data['quitdate'],'%Y-%m-%d %H:%M:%S')
        quitDate = quitDateTime.date()
        timeSnusFree = {
            'years': str(relativedelta(datetime.now(), quitDateTime).years),
            'months': str(relativedelta(datetime.now(), quitDateTime).months),
            'days': str(relativedelta(datetime.now(), quitDateTime).days),
            'hours': str(relativedelta(datetime.now(), quitDateTime).hours),
            'minutes': str(relativedelta(datetime.now(), quitDateTime).minutes)
        }

        prillorSaved = (datetime.now() - quitDateTime).days * data['prillorperday']
        dosorSaved = prillorSaved / data['prillorperdosa']
        moneySaved = dosorSaved * data['priceperdosa']
    else:
        timeSnusFree = None
        quitDate = None
        prillorSaved = None
        dosorSaved = None
        moneySaved = None

    return render(request, 'home.html', {'TimeSnusFree': timeSnusFree, 'QuitDate': quitDate, 'prillorSaved': prillorSaved, 'dosorSaved': dosorSaved, 'moneySaved': moneySaved})

@login_required
def updateView(request) :
    if request.method == "POST":
        requestdata = request.POST.dict()
        logging.debug(str(requestdata))
        # dosPrice = int(''.join(data['pricePerDosa'].split(',')))
        postData = {'quitdate': str(requestdata['quitDate']), 'prillorperday': int(requestdata['prillorPerDay']), 'prillorperdosa': int(requestdata['prillorPerDosa']), 'priceperdosa': int(''.join(requestdata['pricePerDosa'].split(',')))}
        logging.debug(postData)
        # call GoLang API to update data in db
        username = request.user.get_username()
        requestUrl = "http://backend:8080/slutasnusa/" + username 
        logging.debug(requestUrl)
        response = requests.post(requestUrl, data=json.dumps(postData))
        logging.debug(response)
        return HttpResponseRedirect('/home/')

def registerView(request):
    return render(request, 'registration/register.html')

@login_required
def settingsView(request):
    return render(request, 'settings.html')
