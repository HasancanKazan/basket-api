namespace MizuCaseStudy.Core.Utilities
{
    public class SuccessDataResult<T> : DataResult<T>
    {
        public SuccessDataResult(T data) : base(data, true)
        {
        }
        public SuccessDataResult(bool success) : base(default, success)
        {
        }

        public SuccessDataResult(T data, string message) : base(data, true, message)
        {
        }
        public SuccessDataResult(string message) : base(default, true, message)
        {
        }
    }
}
