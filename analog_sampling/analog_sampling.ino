#define LEDPIN 2

void setup() {
  Serial.begin(9600);
  pinMode(LEDPIN, OUTPUT);
}

void loop() {
  digitalWrite(LEDPIN, HIGH);
  int value = analogRead(A0);
  Serial.println(value);
  delay(500);
  digitalWrite(LEDPIN, LOW);
  value = analogRead(A0);
  Serial.println(value);
  delay(500);
}
