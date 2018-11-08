import java.util.*;
import java.io.*;
import java.math.*;

public class FactorilaVsExponential {
    public static void main(String[] args) {

        Scanner in = new Scanner(System.in);
        int K = in.nextInt();
        double[] values = new double[K];
        double valueTmp;
        int factorielTmp;
        int puissanceTmp;

        for (int i = 0; i < K; i++) {
            float A = in.nextFloat();
            values[i] = (double) A;
        }

        for (int i = 0; i < values.length; i++) {
            valueTmp = values[i];

            System.err.println("Traitement : " + valueTmp);

            int cpt = 0;
            while (cpt != -1) {

                factorielTmp = calculFactoriel(cpt);
                puissanceTmp = calculPuissance(valueTmp, (double) cpt);

                System.err.println("\t" + puissanceTmp + " < " + factorielTmp);
                if (puissanceTmp < factorielTmp) {
                    System.out.print(cpt);
                    break;
                }
                cpt++;
            }
            System.out.print(" ");
        }
    }

    public static int calculPuissance(double value, double puissance) {
        int sum = 0;

        sum = (int) Math.pow(value, puissance);

        return sum;
    }

    public static int calculFactoriel(int value) {
        int result = 1;

        if (value > 0) {
            for (int i = 1; i <= value; i++) {
                result *= value;
            }
        }

        System.err.print("\ncalculFactoriel " + value + " ! ");
        System.err.println("\t==> " + result);

        return result;
    }
}