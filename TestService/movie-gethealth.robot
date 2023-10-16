*** Settings ***
Variables         ./resources/movie-gethealth.yaml

Library           RequestsLibrary
Library           DatabaseLibrary
Library           urllib3
Library           DateTime
Library           String
Library           Collections
Library           DateTime
Library           JSONLibrary
Library           OperatingSystem
# Library           ../lib/common.py

*** Variable ***
${BACKEND_HOST}     http://localhost:8080
${API_KEY}          1234567890
${uri_movie}        /project-os-container/movie/

*** Keywords ***
Call api for health
    [Arguments]
    ${headers}=     Create Dictionary   Api_Key=${API_KEY}
    Create Session     get_health      ${BACKEND_HOST}
    ${response}=      GET On Session      get_health    ${uri_movie}    headers=${headers}     expected_status=any
    Set Test Variable    ${response}
    Log    ${response.text}
    [Return]    ${response}

Call api for health without api key
    [Arguments]
    ${headers}=     Create Dictionary    Api-Key=null
    ${response}=      GET On Session      get_health    ${uri_movie}    headers=${headers}     expected_status=any
    Set Test Variable    ${response}
    Log    ${response.text}
    [Return]    ${response}

Verify response http status
    [Arguments]     ${expect}
    Should Be Equal As Integers     ${response.status_code}     ${expect}

Verify health response body
    [Arguments]     ${expect}
    &{json}=    Set Variable        ${response.json()}
    Should Be Equal As Strings      ${expect.code}   ${json.code}
    Should Be Equal As Strings      ${expect.status}   ${json.status}
    Should Be Equal As Strings      ${expect.description}   ${json.description}
    Should Be Equal As Strings      ${expect.message}    ${json.message}

Verify error response body
    [Arguments]     ${expect}
    &{json}=    Set Variable        ${response.json()}
    Should Be Equal As Strings      ${expect.code}      ${json.code}
    Should Be Equal As Strings      ${expect.error}     ${json.error}
    Should Be Equal As Strings      ${expect.thai_description}      ${json.thai_description}
    Should Be Equal As Strings      ${expect.english_description}   ${json.english_description}

*** Test Cases ***
TC_BE_00001 success - get health from news service found
    [Documentation]    to verify how api return health from news service
    [Tags]    ready    success    regression
    When Call api for health
    Then Verify response http status  200
    And Verify health response body  ${success.expect}

TC_BE_00002 fail - get health from news service without api key
    [Documentation]    to verify how api return health from news service without api key
    [Tags]    ready    fail    regression
    When Call api for health without api key
    Then Verify response http status  401
    And Verify error response body  ${fail.missing_api_key}