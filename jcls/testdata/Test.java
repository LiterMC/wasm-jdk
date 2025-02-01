
public class Test {
	private int privateInt0 = 0;
	int int1 = 1;
	public long publicLong16 = 16;
	public final String publicFinalString = "publicFinalString";

	private void testPrivateVoidMethod() {
		System.out.println("testPrivateVoidMethod");
	}

	public final long testPublicFinalAddMethod(int a, long b) {
		return a + b + this.int1;
	}

	public static void main(String[] args) {
		if (args.length > 0) {
			System.out.println("arg0: " + args[0]);
		}
		System.out.print("Test class " + args.length + "\n");
		Test tt = new Test();
		System.out.println("running testPrivateVoidMethod");
		tt.testPrivateVoidMethod();
		System.out.println("running testPublicFinalAddMethod");
		long x = tt.testPublicFinalAddMethod(3, 2);
		System.out.println(x);
	}
}
