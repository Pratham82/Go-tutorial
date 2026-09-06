/* Reusable quiz widget for Go-job-prep lessons.
 *
 * Usage in a lesson:
 *   <div class="quiz"
 *        data-question="What does %w do in fmt.Errorf?"
 *        data-options='["Wraps the error so it can be unwrapped","Formats the error as a string only","Discards the wrapped error","Panics if the error is nil"]'
 *        data-answer="0"
 *        data-explain="%w records the wrapped error so errors.Is / errors.As can find it later."></div>
 *
 * Multiple .quiz blocks per page are fine. Feedback is immediate and automatic.
 * Keep every option the same length where possible so formatting gives nothing away.
 */
(function () {
  function build(el) {
    var q = el.getAttribute("data-question");
    var opts = JSON.parse(el.getAttribute("data-options") || "[]");
    var answer = parseInt(el.getAttribute("data-answer"), 10);
    var explain = el.getAttribute("data-explain") || "";

    var qEl = document.createElement("div");
    qEl.className = "quiz-q";
    qEl.textContent = q;
    el.appendChild(qEl);

    var fb = document.createElement("div");
    fb.className = "quiz-fb";

    var buttons = [];
    opts.forEach(function (text, i) {
      var b = document.createElement("button");
      b.className = "quiz-opt";
      b.type = "button";
      b.textContent = text;
      b.addEventListener("click", function () {
        var correct = i === answer;
        buttons.forEach(function (bb, j) {
          bb.disabled = true;
          if (j === answer) bb.classList.add("correct");
          if (j === i && !correct) bb.classList.add("wrong");
        });
        fb.classList.add(correct ? "correct" : "wrong");
        fb.textContent = (correct ? "Correct. " : "Not quite. ") + explain;
      });
      buttons.push(b);
      el.appendChild(b);
    });

    el.appendChild(fb);
  }

  document.addEventListener("DOMContentLoaded", function () {
    document.querySelectorAll(".quiz").forEach(build);
  });
})();
